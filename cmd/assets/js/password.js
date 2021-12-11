;let vm = new Vue({
    delimiters: ['${', '}'],
    el: '#app',
    data: {
        form: {
            password: '',
            confirm: ''
        },
        loading: false
    },
    methods: {
        submit() {
            this.loading = true;

            axios.post('/password/change', this.form).then(function (response) {
                vm.loading = false;
                let resp = response.data;

                if (!resp.err) {
                    vm.$message({
                        type: 'success',
                        message: resp.msg,
                        duration: 1000,
                        onClose: function() {
                            location.href = '/';
                        }
                    });
                } else {
                    vm.$message.error(resp.msg);
                }
            }).catch(function (err) {
                vm.loading = false;
                vm.$message.error(resp.msg);

                console.log('哎呀！服务器开了点小差')
            });
        },
        cancel() {
            window.history.back(-1);
        }
    }
});