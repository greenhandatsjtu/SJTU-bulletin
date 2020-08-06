<template>
  <v-app v-scroll="onScroll">
    <v-app-bar
        app
        color="indigo"
        dark
    >
      <div class="d-flex align-center">
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

    <v-main>
      <v-row class="px-2 pt-1 ma-auto">
        <v-chip
            class="mx-2 mt-1 font-weight-bold"
            color="primary"
            text-color="white"
        >
          <v-avatar left>
            <v-icon>mdi-account-circle</v-icon>
          </v-avatar>
          Visitor {{visitor}}
        </v-chip>
        <v-chip
            class="mx-2 mt-1 font-weight-bold"
            color="green"
            text-color="white"
        >
          <v-avatar left>
            <v-icon>mdi-web</v-icon>
          </v-avatar>
          Request {{request}}
        </v-chip>
      </v-row>
      <Bulletin/>
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
  }),
  methods: {
    onScroll(e) {
      this.offsetTop = document.scrollingElement.scrollTop
    },
  },
  beforeCreate() {
    this.$axios.get('/visit').then(({data}) => {
      this.visitor = data.visitor
      this.request = data.request
    })
  }
};
</script>
