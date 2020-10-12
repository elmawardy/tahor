<template>
    <div style="min-height:80vh">
        <p class="title">New User</p>
        <div class="columns">
            <div class="column is-half">
                <section class="modal-card-body">
                    <b-field label="Email">
                        <b-input
                            type="email"
                            v-model="email"
                            placeholder="example@tahor.com"
                            required>
                        </b-input>
                    </b-field>
                    <b-field label="Name">
                        <b-input
                            type="text"
                            v-model="name"
                            placeholder="John Doe">
                        </b-input>
                    </b-field>
                    <b-field label="Password">
                        <b-input
                            type="password"
                            v-model="password"
                            password-reveal>
                        </b-input>
                    </b-field>
                    <b-field label="Retype Password">
                        <b-input
                            type="password"
                            v-model="retypePassword"
                            password-reveal
                            v-on:keyup.native.enter="register">
                        </b-input>
                    </b-field>

                    <br>
                    <button :loading="loading" @click="register" :class="`button is-primary ${loading ? 'is-loading':''}`">Create</button>
                </section>
            </div>
        </div>
    </div>
</template>

<script>
import { ToastProgrammatic as Toast } from "buefy";
export default {
    data() {
        return {
            email:'',
            name:'',
            password:'',
            retypePassword:'',
            loading:false
        }
    },
    methods: {
        register(){
            this.loading = true
            fetch(this.$store.state.authURL + "/api/auth/register",{
                headers: {
                Accept: "application/json",
                    "Content-Type": "application/json"
                },
                method: "POST",
                body: JSON.stringify({
                    email:this.email,
                    password:this.password,
                    name:this.name
                })
            })
            .then(response => response.json())
            .then(response => {
                if (response.name){
                    Toast.open(response.message);
                }else{
                    console.log(response)
                    Toast.open("Done");
                }
                this.loading = false
            })
            .catch(err => {
                console.log(err)
                Toast.open(err);
            })
        }
    },
}
</script>