<template>
    <div style="min-height:80vh">
        <p class="title">Add case</p>
        <div class="columns">
            <div class="column is-half">
                <section class="modal-card-body">
                    <b-field label="Description">
                        <b-input maxlength="200" type="textarea"></b-input>
                    </b-field>
                    <b-field label="Target">
                        <b-input
                            type="number"
                            v-model="target">
                        </b-input>
                    </b-field>
                    <b-field label="Collected">
                        <b-input
                            type="number"
                            v-model="collected">
                        </b-input>
                    </b-field>
                    <b-field label="Currency">
                        <b-select placeholder="Select a character" disabled>
                            <option value="flint">EGP</option>
                            <option value="silver">PLN</option>
                            <option value="vane">UK Pound</option>
                        </b-select>
                    </b-field>
                    <b-field label="Enter some tags">
                        <b-taginput
                            v-model="tags"
                            :data="filteredTags"
                            autocomplete
                            :allow-new="true"
                            :open-on-focus="true"
                            field="tag"
                            icon="label"
                            placeholder="Add a tag"
                            @typing="getFilteredTags">
                        </b-taginput>
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
            description:'',
            target:null,
            collected:null,
            loading:false,
            tags: ['Health'],
            filteredTags: ['Food','Water'],
        }
    },
    methods: {
        getFilteredTags(text) {
            this.filteredTags = this.tags.filter((option) => {
                return option
                    .toString()
                    .toLowerCase()
                    .indexOf(text.toLowerCase()) >= 0
            })
        },
        register(){
            this.loading = true
            fetch(this.$store.state.caseURL + "/api/cases/add",{
                headers: {
                Accept: "application/json",
                    "Content-Type": "application/json"
                },
                method: "POST",
                body: JSON.stringify({
                    desc:this.description,
                    target:this.target,
                    collected:this.collected,
                    tags:this.tags,
                    currency:this.currency
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