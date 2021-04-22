<template>
  <v-app v-scroll="onScroll">
    <v-app-bar
        app
        color="#004098"
        dark
        clipped-left
    >
      <div class="d-flex align-center">
        <v-app-bar-nav-icon @click.stop="drawer = !drawer"></v-app-bar-nav-icon>
        <v-img
            alt="SJTU Logo"
            class="shrink mr-2"
            contain
            :src="require('@/assets/logo.png')"
            transition="scale-transition"
            width="40"
        />
        <v-toolbar-title class="font-italic font-weight-bold">SJTU Bulletin</v-toolbar-title>
      </div>

      <v-spacer></v-spacer>

      <v-btn
          href="https://github.com/greenhandatsjtu/SJTU-bulletin"
          target="_blank"
          text
      >
        <span class="mr-2">GitHub</span>
        <v-icon>mdi-open-in-new</v-icon>
      </v-btn>
    </v-app-bar>

    <v-navigation-drawer v-model="drawer" app clipped>
      <v-card color="#004098" dark flat tile>
        <v-list-item>
          <v-list-item-content>
            <v-list-item-title class="title pt-2 text-center">
              Preference
            </v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-card>

      <v-list dense nav>
        <v-list-item :key="index" v-for="(checkbox, index) in checkboxes">
          <v-checkbox :color="checkbox.color" :label="checkbox.label" :value="checkbox.value" class="pl-4 my-2"
                      hide-details
                      v-model="selected"></v-checkbox>
        </v-list-item>
      </v-list>

      <template v-slot:append>
        <v-row class="px-0 pb-2 ma-auto">
          <v-chip
              class="mx-2 mt-0 mb-2 font-weight-bold"
              color="primary"
              text-color="white"
          >
            <v-avatar left>
              <v-icon>mdi-account-circle</v-icon>
            </v-avatar>
            Visit {{ visitor }}
          </v-chip>
          <v-chip
              class="mx-2 mt-0 mb-2 font-weight-bold"
              color="green"
              text-color="white"
          >
            <v-avatar left>
              <v-icon>mdi-web</v-icon>
            </v-avatar>
            Request {{ request }}
          </v-chip>
        </v-row>
        <v-footer
            class="justify-center"
        >
          <span>greenhandatsjtu &commat; {{ new Date().getFullYear() }}</span>
        </v-footer>
      </template>
    </v-navigation-drawer>

    <v-main>
      <Bulletin :selected="selected"/>
    </v-main>
    <v-fab-transition mode="">
      <v-btn
          v-show="offsetTop>=1000"
          @click="$vuetify.goTo(0)"
          bottom
          large
          color="pink lighten-3"
          dark
          fab
          fixed
          right
      >
        <v-icon>mdi-arrow-up</v-icon>
      </v-btn>
    </v-fab-transition>
  </v-app>
</template>

<script>
import Bulletin from "./views/Bulletin";

export default {
  name: 'App',

  components: {
    Bulletin
  },

  data: () => ({
    offsetTop: 0,
    visitor: 0,
    request: 0,
    drawer: null,
    selected: [],
    checkboxes: [
      {value: 'jwc', label: '教务处', color: 'blue'},
      {value: 'xsb', label: '学生办', color: 'teal'},
      {value: 'sjtuNotice', label: '通知通告', color: 'red'},
      {value: 'partTime', label: '实习信息', color: 'green'},
      {value: 'fullTime', label: '全职招聘', color: 'orange'},
      {value: 'ourHome', label: '生活园区', color: 'cyan'},
    ]
  }),
  methods: {
    onScroll(e) {
      this.offsetTop = document.scrollingElement.scrollTop
    },
  },
  watch: {
    selected: function (newVal) {
      localStorage.setItem('selection', JSON.stringify(newVal))
    }
  },
  beforeCreate() {
    this.$axios.get('/visit').then(({data}) => {
      this.visitor = data.visitor
      this.request = data.request
    })
  },
  created() {
    let selected = JSON.parse(localStorage.getItem('selection'))
    if (selected == null) {
      selected = []
    }
    this.selected = selected
  }
};
</script>
