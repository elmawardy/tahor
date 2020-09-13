<template>
  <div>
    <div class="modal-card" style="width: auto">
      <header class="modal-card-head">
        <div class="modal-card-title">
          <div class="tabs is-centered">
            <ul>
              <li :class="selectedTab == 'login'? 'is-active' :''" @click="selectedTab='login'">
                <a>Login</a>
              </li>
              <li
                :class="selectedTab == 'register'? 'is-active' :''"
                @click="selectedTab='register'"
              >
                <a>Register</a>
              </li>
            </ul>
          </div>
        </div>
      </header>
      <div
        :style="{visibility: selectedTab == 'register' ? 'visible' : 'hidden'}"
        v-if="selectedTab == 'register'"
      >
        <section class="modal-card-body">
          <b-field label="Name">
            <b-input type="text" v-model="name" placeholder="Your Name" required></b-input>
          </b-field>

          <b-field label="Email">
            <b-input type="email" v-model="registerEmail" placeholder="Your email" required></b-input>
          </b-field>

          <b-field label="Password">
            <b-input
              type="password"
              v-model="registerpassword"
              password-reveal
              placeholder="Your password"
              required
            ></b-input>
          </b-field>

          <b-field label="Password Retype">
            <b-input
              type="password"
              v-model="passwordretype"
              password-reveal
              placeholder="Your password again"
              required
            ></b-input>
          </b-field>
        </section>
        <footer class="modal-card-foot">
          <button class="button" type="button" @click="$parent.close()">غلق</button>
          <button
            :class="`button is-primary ${loadingRegister ? 'is-loading':''}`"
            @click="register()"
          >تسجيل</button>
        </footer>
      </div>
      <div v-if="selectedTab == 'login'">
        <section class="modal-card-body">
          <b-field label="Email">
            <b-input type="email" v-model="email" placeholder="Your email" required></b-input>
          </b-field>

          <b-field label="Password">
            <b-input
              type="password"
              v-model="password"
              password-reveal
              placeholder="Your password"
              required
            ></b-input>
          </b-field>

          <b-checkbox>Remember me</b-checkbox>
        </section>
        <footer class="modal-card-foot">
          <button class="button" type="button" @click="$parent.close()">Close</button>
          <button
            @click="login"
            :class="`button is-primary ${loadingLogin ? 'is-loading':''}`"
          >Login</button>
        </footer>
      </div>
    </div>
  </div>
</template>

<script>
import { ToastProgrammatic as Toast } from "buefy";
export default {
  props: ["login_or_register"],
  data() {
    return {
      loadingRegister: false,
      loadingLogin: false,
      name: "",
      email: "",
      registerEmail: "",
      password: "",
      registerpassword: "",
      passwordretype: "",
      selectedTab: ""
    };
  },
  mounted() {
    this.selectedTab = this.login_or_register;
  },
  methods: {
    login() {
      this.loadingLogin = true;
      fetch(this.$store.state.backendURL + "/api/user/login", {
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json"
        },
        method: "POST",
        body: JSON.stringify({
          email: this.email,
          password: this.password
        })
      })
        .then(response => response.json())
        .then(response => {
          if (response.State == "Success") {
            window.localStorage.setItem("jwt", response.jwt);
            fetch(this.$store.state.backendURL + "/api/user/getbasicinfo", {
              headers: {
                Accept: "application/json",
                "Content-Type": "application/json"
              },
              method: "POST",
              body: JSON.stringify({
                jwt: window.localStorage.getItem("jwt")
              })
            })
              .then(response => response.json())
              .then(response => {
                if (response.State == "Success") {
                  window.localStorage.setItem(
                    "user_name",
                    response.user_info.Name
                  );
                  this.loadingLogin = false;
                  Toast.open("مرحبا");
                  this.$store.commit("login");
                  this.$store.commit("setUsername", response.user_info.Name);
                  this.$parent.close();
                }
              });
          } else {
            this.loadingLogin = false;
            Toast.open("خطأ");
          }
        });
    },
    register() {
      this.loadingRegister = true;
      var that = this;

      fetch(this.$store.state.backendURL + "/api/user/register", {
        headers: {
          Accept: "application/json",
          "Content-Type": "application/json"
        },
        method: "POST",
        body: JSON.stringify({
          email: this.registerEmail,
          name: this.name,
          password: this.registerpassword
        })
      })
        .then(response => response.json())
        .then(response => {
          this.loadingRegister = false;
          if (response.State == "Success") {
            Toast.open("مرحبا" + this.name);
            that.$parent.close();
          } else {
            Toast.open("خطأ");
          }
        });
    }
  }
};
</script>