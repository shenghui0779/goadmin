;let vm = new Vue({
    delimiters: ['${', '}'],
    el: '#app',
    data: {
        collapse: false,
        unfreeze: [],
        loading: {
            app: false,
            search: false
        }
    },
    methods: {
        init() {
            this.$notify.success({
                title: '欢迎登录',
                message: 'welcome to goadmin'
            });

            console.log('home')
        },
        menuToggle() {
            this.collapse = !this.collapse;
        }
    }
});

vm.init();