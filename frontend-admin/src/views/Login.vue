<template>
  <div class="columns">
    <div class="column is-one-third"></div>
    <div class="column is-one-third">
        <div class="modal-card" style="width: auto">
        <header class="modal-card-head">
            <div class="modal-card-title">
            Signin
            </div>
        </header>
        <div>
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
            <button
                @click="login"
                :class="`button is-primary ${loadingLogin ? 'is-loading':''}`"
            >Login</button>
            </footer>
        </div>
        </div>
    </div>
  </div>
</template>

<script>
import { ToastProgrammatic as Toast } from "buefy";
export default {
  data() {
    return {
      loadingLogin: false,
      name: "",
      email: "",
      password: "",
    };
  },
  methods: {
    login() {
      this.loadingLogin = true;
      fetch(this.$store.state.authURL + "/api/auth/signin", {
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
        .then(res => {
          if(res.ok) {
            return res.json()
          }
          return res.json().then(json => {
            this.loadingLogin = false;
            Toast.open(json.message);
          })
        })
        .then(response => {
            window.localStorage.setItem("jwt", response.jwt);
            fetch(this.$store.state.authURL + "/api/auth/getbasicinfo", {
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
                  window.localStorage.setItem(
                    "user_name",
                    response.name
                  );
                  this.loadingLogin = false;
                  Toast.open("مرحبا");
                  this.$store.commit("login");
                  this.$store.commit("setUsername", response.name);
              });
        })
        .catch(err => {
            console.log(err)
        });
    },
  }
};
</script>