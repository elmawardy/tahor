<template>
  <div class="container is-fluid">
    <div class="content-container" :style="`margin-left:${menuCollapsed ? '50px;' : '200px'}`">
      <b-navbar style="z-index:9;">
        <template slot="brand">
          <b-navbar-item tag="router-link" :to="{ path: '/' }">
            <img
              src="https://raw.githubusercontent.com/buefy/buefy/dev/static/img/buefy-logo.png"
              alt="Lightweight UI components for Vue.js based on Bulma"
            />
          </b-navbar-item>
        </template>
        <template slot="start">
          <!-- <b-navbar-item href="/">الصفحة الرئيسية</b-navbar-item> -->
        </template>

        <template slot="end" v-if="$store.state.loggedIn">
          <div class="navbar-item has-dropdown is-hoverable">
            <a class="navbar-link">مرحبا {{$store.state.user_name}}</a>

            <div class="navbar-dropdown">
              <a class="navbar-item">الملف الشخصي</a>
              <a class="navbar-item">الحالات المحفوظة</a>
              <hr class="navbar-divider" />
              <a class="navbar-item" @click="Logout">تسجيل خروج</a>
            </div>
          </div>
        </template>
      </b-navbar>

      <router-view/>
      

      <footer class="footer">
        <div class="content has-text-centered">
          <p>
            <strong>Bulma</strong> by
            <a href="https://jgthms.com">Jeremy Thomas</a>. The source code is licensed
            <a
              href="http://opensource.org/licenses/mit-license.php"
            >MIT</a>. The website content
            is licensed
            <a
              href="http://creativecommons.org/licenses/by-nc-sa/4.0/"
            >CC BY NC SA 4.0</a>.
          </p>
        </div>
      </footer>
    </div>
    <sidebar-menu
      :menu="menu"
      :width="'200px'"
      :theme="'white-theme'"
      style="z-index:10"
      :collapsed="menuCollapsed"
      @toggle-collapse="onToggleCollapse"
    />
  </div>
</template>

<script>

export default {
  data() {
    return {
      menuCollapsed: false,
      menu: [
        {
          href: "/",
          title: "الصفحة الرئيسية",
          icon: "fa fa-home"
        },
        {
          href: "/cases",
          title: "الحالات",
          icon: "fa fa-heartbeat"
        },
        {
          href: "/users",
          title: "المستخدمين",
          icon: "fa fa-users"
        },
      ]
    };
  },
  methods:{
     onToggleCollapse(collapsed) {
      this.menuCollapsed = collapsed;
    },
  }
};
</script>

<style lang="scss">
@import "~bulma/sass/utilities/_all";
@import "~bulma-divider";
@import "~bulma-tooltip";

// Set your colors
$primary: #3a9e5d;
$primary-invert: findColorInvert($primary);
$twitter: #4099ff;
$twitter-invert: findColorInvert($twitter);

// Setup $colors to use as bulma classes (e.g. 'is-twitter')
$colors: (
  "white": (
    $white,
    $black
  ),
  "black": (
    $black,
    $white
  ),
  "light": (
    $light,
    $light-invert
  ),
  "dark": (
    $dark,
    $dark-invert
  ),
  "primary": (
    $primary,
    $primary-invert
  ),
  "info": (
    $info,
    $info-invert
  ),
  "success": (
    $success,
    $success-invert
  ),
  "warning": (
    $warning,
    $warning-invert
  ),
  "danger": (
    $danger,
    $danger-invert
  ),
  "twitter": (
    $twitter,
    $twitter-invert
  )
);

// Links
$link: $primary;
$link-invert: $primary-invert;
$link-focus-border: $primary;

// Import Bulma and Buefy styles
@import "~bulma";
@import "~buefy/src/scss/buefy";

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;
}

#nav a {
  font-weight: bold;
  color: #2c3e50;
}

#nav a.router-link-exact-active {
  color: #42b983;
}

.v-sidebar-menu .vsm--link.vsm--link_active {
  .vsm--icon {
    background-color: $primary !important;
  }
}

.v-sidebar-menu .vsm--badge {
  background-color: $primary !important;
  color: $primary;
}

.content-container {
  transition: margin 0.4s;
}
</style>
