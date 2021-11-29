;let vm = new Vue({
    delimiters: ['${', '}'],
    el: '#app',
    data: {
        collapse: false,
        list: {
            count: 0,
            data: [],
        },
        roles: [
            {
                "id": 1,
                "name": "超级管理员",
            },
            {
                "id": 2,
                "name": "高级管理员",
            },
            {
                "id": 3,
                "name": "普通管理员",
            },
            {
                "id": 4,
                "name": "微博",
            },
        ],
        user: {},
        curPage: 1,
        search: {
            page: 1,
            size: 10,
            name: '',
            role: ''
        },
        dialog: {
            add: false,
            edit: false
        },
        loading: {
            app: false,
            search: false
        }
    },
    methods: {
        init() {},
        menuToggle() {
            this.collapse = !this.collapse;
        },
        query(curPage) {
            this.loading.search = true;
            this.curPage = curPage;
            this.search.page = curPage;

            axios.post('/users/query', {
                page: this.search.page,
                size: this.search.size,
                name: this.search.name,
                role: this.search.role || 0,
            }).then(function(response) {
                vm.loading.search = false;
                let resp = response.data;

                if (!resp.err) {
                    if (vm.search.page == 1) {
                        vm.list.count = resp.data.count;
                    }

                    vm.list.data = resp.data.list;
                } else {
                    vm.$message.error(resp.msg);
                }
            }).catch(function(err) {
                vm.loading.search = false;
                vm.$message.error('哎呀！服务器开小差了');

                console.log(err);
            });
        },
        sizeChange(size) {
            this.search.size = size;
            this.query(1);
        },
        add() {
            this.user = {};
            this.dialog.add = true;
        },
        edit(index, row) {
            this.user = row;
            this.dialog.edit = true;
        },
        submit() {
            this.loading.app = true;

            axios.post('/users/add', {
                name: this.user.name,
                email: this.user.email,
                role_id: this.user.role || 0
            }).then(function(response) {
                vm.loading.app = false;
                let resp = response.data;

                if (!resp.err) {
                    vm.dialog.add = false
                    vm.$message.success(resp.msg);
                    vm.query(1);
                } else {
                    vm.$message.error(resp.msg);
                }
            }).catch(function(err) {
                vm.loading.app = false;
                vm.$message.error('哎呀！服务器开小差了');

                console.log(err);
            });
        },
        save() {
            this.loading.app = true;

            axios.post('/users/edit', {
                id: this.user.id,
                name: this.user.name,
                email: this.user.email,
                role: this.user.role || 0,
            }).then(function(response) {
                vm.loading.app = false;
                let resp = response.data;

                if (!resp.err) {
                    vm.dialog.edit = false
                    vm.$message.success(resp.msg);
                    vm.query(vm.curPage);
                } else {
                    vm.$message.error(resp.msg);
                }
            }).catch(function (err) {
                vm.loading.app = false;
                vm.$message.error('哎呀！服务器开小差了');

                console.log(err);
            });
        },
        reset(index, row) {
            this.$confirm('确定重置密码?', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(function() {
                vm.loading.app = true;

                axios.post('/password/reset', {id: row.id}).then(function(response) {
                    vm.loading.app = false;
                    let resp = response.data;

                    if (!resp.err) {
                        vm.$message.success(resp.msg);
                    } else {
                        vm.$message.error(resp.msg);
                    }
                }).catch(function(err) {
                    vm.loading.app = false;
                    vm.$message.error('哎呀！服务器开小差了');

                    console.log(err);
                });
            }).catch(function() {
                console.log('reset cancel');
            });
        },
        del(index, row) {
            this.$confirm('确定删除？', '提示', {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                type: 'warning'
            }).then(function() {
                vm.loading.app = true;

                axios.post('/users/delete', {id: row.id}).then(function(response) {
                    vm.loading.app = false;
                    let resp = response.data;

                    if (!resp.err) {
                        vm.$message.success(resp.msg);
                        vm.query(1);
                    } else {
                        vm.$message.error(resp.msg);
                    }
                }).catch(function(err) {
                    vm.loading.app = false;
                    vm.$message.error('哎呀！服务器开小差了');

                    console.log(err);
                });
            }).catch(function() {
                console.log('delete cancel');
            });
        }
    }
});

vm.init();
vm.query(1);