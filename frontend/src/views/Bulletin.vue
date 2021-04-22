<template>
  <v-container class="px-4" fluid>
    <v-snackbar
        :color="snackbar.status"
        :timeout="2000"
        app
        rounded
        top
        v-model="snackbar.open"
    >
      {{ snackbar.message }}
      <template v-slot:action="{ attrs }">
        <v-btn
            @click="snackbar.open = false"
            icon
        >
          <v-icon>mdi-close</v-icon>
        </v-btn>
      </template>
    </v-snackbar>
    <v-list>
      <!--            使用自定义noticeItem组件-->
      <notice-item :key="index" :notice="notice" v-for="(notice,index) in showedNotices"></notice-item>
    </v-list>
    <infinite-loading :identifier="infiniteId" @infinite="infiniteHandler"></infinite-loading>
  </v-container>
</template>

<script>
import NoticeItem from "../components/NoticeItem";
import InfiniteLoading from 'vue-infinite-loading';

export default {
  name: "Bulletin",
  props: ['selected'],
  components: {NoticeItem, InfiniteLoading},
  data: () => ({
    page: 0,
    infiniteId: +new Date(),
    notices: [],
    showedNotices: [],
    snackbar: {
      open: false,
      message: null,
      status: "success",
    },
  }),
  methods: {
    infiniteHandler($state) {
      this.$axios.get('/notices', {
        params: {
          page: this.page
        }
      }).then(({data}) => {
        if (data.length === 0) {
          $state.complete();
          return
        }
        this.page += 1
        this.notices.push(...data)
        if (this.selected.length === 0) {
          this.showedNotices.push(...data)
        } else {
          for (let i = 0; i < data.length; i++) {
            if (this.selected.indexOf(data[i]['_type']) !== -1) {
              this.showedNotices.push(data[i])
            }
          }
        }
        $state.loaded();
      }).catch(() => {
        this.snackbar.message = "Fetch notices error"
        this.snackbar.status = "error"
        this.snackbar.open = true
        $state.complete()
      })
    },
  },
  watch: {
    // 当复选框有变化，展示的通知要进行相应筛选
    selected: function (newVal) {
      this.infiniteId += 1;
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
}
</script>

<style scoped>

</style>
