; let vm = new Vue({
    delimiters: ['${', '}'],
    el: '#app',
    data: {
        captcha: '',
        login: {
            id: '',
            account: '',
            password: '',
            captcha: ''
        },
        loading: false
    },
    methods: {
        refreshCaptcha() {
            axios.get('/captcha').then(function (response) {
                let resp = response.data;

                if (!resp.err) {
                    vm.login.id = resp.data.id;
                    vm.captcha = resp.data.content;
                } else {
                    vm.$message.error(resp.msg);
                }
            }).catch(function (err) {
                vm.$message.error('哎呀！服务器开了点小差');

                console.log(err)
            });
        },
        submit() {
            this.loading = true;

            axios.post('/login', this.login).then(function (response) {
                vm.loading = false;
                let resp = response.data;

                if (!resp.err) {
                    vm.$message({
                        type: 'success',
                        message: resp.msg,
                        duration: 1000,
                        onClose: function () {
                            location.href = '/';
                        }
                    });
                } else {
                    vm.refreshCaptcha();
                    vm.$message.error(resp.msg);
                }
            }).catch(function (err) {
                vm.refreshCaptcha();
                vm.$message.error('哎呀！服务器开了点小差');

                console.log(err)
            });
        }
    }
});

document.onkeydown = function (e) {
    if (e.keyCode === 13) {
        vm.submit();
    }
};

vm.refreshCaptcha();