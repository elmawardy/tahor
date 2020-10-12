<template>
<div class="modal-card" style="width: auto">
    <header class="modal-card-head">
        <p class="modal-card-title">Edit Case #{{case_id}}</p>
    </header>
    <section class="modal-card-body">
        <b-field label="Severity">
            <b-input
                type="number"
                :value="email"
                placeholder="1-10"
                required>
            </b-input>
        </b-field>
         <b-field label="Category">
            <b-input
                type="text"
                :value="email"
                placeholder="Category"
                required>
            </b-input>
        </b-field>
        <b-field label="Description">
            <b-input maxlength="200" type="textarea"></b-input>
        </b-field>
        <b-field label="Date Modified">
            <b-datepicker
                v-model="datemodified"
                placeholder="Click to select..."
                icon="calendar-today"
                inline
                trap-focus>
            </b-datepicker>
        </b-field>


    </section>
    <footer class="modal-card-foot">
        <button class="button" type="button" @click="$parent.close()">Cancel</button>
        <button class="button is-primary">Save</button>
    </footer>
</div>
</template>

<script>
export default {
    props: ['email', 'password','case_id'],
    data() {
        return {
            datemodified: new Date('10/8/2020')
        }
    },
}
</script>