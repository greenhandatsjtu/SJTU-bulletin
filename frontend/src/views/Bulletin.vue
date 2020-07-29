<template>
    <v-container fluid class="px-4">
        <v-snackbar
                rounded
                app
                v-model="snackbar.open"
                :color="snackbar.status"
                :timeout="2000"
                top
        >
            {{ snackbar.message }}
            <template v-slot:action="{ attrs }">
                <v-btn
                        icon
                        @click="snackbar.open = false"
                >
                    <v-icon>mdi-close</v-icon>
                </v-btn>
            </template>
        </v-snackbar>
        <v-divider></v-divider>
        <v-card class="mt-2">
            <!--            复选框-->
            <v-row dense justify="center" align="center">
                <v-col v-for="(checkbox, index) in checkboxes" :key="index">
                    <v-checkbox hide-details :color="checkbox.color" class="px-0 my-2 mx-auto" v-model="selected" :label="checkbox.label"
                                :value="checkbox.value"></v-checkbox>
                </v-col>
            </v-row>
        </v-card>
        <v-list>
            <!--            加载效果-->
            <v-skeleton-loader
                    v-show="this.showedNotices.length===0"
                    ref="skeleton"
                    type="article"
                    class="mx-auto"
            ></v-skeleton-loader>
            <!--            使用自定义noticeItem组件-->
            <notice-item v-for="(notice,index) in showedNotices" :key="index" :notice="notice"></notice-item>
        </v-list>
    </v-container>
</template>

<script>
    import NoticeItem from "../components/NoticeItem";

    export default {
        name: "Bulletin",
        components: {NoticeItem},
        data: () => ({
            notices: [],
            showedNotices: [],
            selected: [],
            snackbar: {
                open: false,
                message: null,
                status: "success",
            },
            checkboxes: [
                {value: 'jwc', label: '教务处', color: 'blue'},
                {value: 'xsb', label: '学生办', color: 'teal'},
                {value: 'sjtuNotice', label: '通知通告', color: 'red'},
                {value: 'partTime', label: '实习信息', color: 'green'},
                {value: 'fullTime', label: '全职招聘', color: 'orange'},
                {value: 'ourHome', label: '生活园区', color: 'grey'},
            ]
        }),
        watch: {
            // 当复选框有变化，展示的通知要进行相应筛选
            selected: function (newVal) {
                if (newVal.length === 0) {
                    this.showedNotices = this.notices
                    return
                }
                this.showedNotices = []
                for (let i = 0; i < this.notices.length; i++) {
                    if (newVal.indexOf(this.notices[i]['_type']) !== -1) {
                        this.showedNotices.push(this.notices[i])
                    }
                }
            }
        },
        created: function () {
            // 从服务器获取消息
            this.$api.get('/notices',)
                .then((response) => {
                    // handle success
                    this.notices = response.data
                    this.showedNotices = response.data
                })
                .catch(() => {
                    // handle error
                    this.snackbar.message = "Fetch notices error"
                    this.snackbar.status = "error"
                    this.snackbar.open = true
                })
        }
    }
</script>

<style scoped>

</style>
