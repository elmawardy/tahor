<template>
    <div style="padding:10px;min-height:100vh">
        <p class="title">Cases Manager</p>
        <b-button 
            icon-left="plus">
            New Case
        </b-button>
        <b-input placeholder="Search..."
            type="search"
            icon="magnify"
            icon-clickable
            style="margin-top:10px;">
        </b-input>
        <b-table 
            :hoverable="true" 
            :paginated="true" 
            :data="data" 
            per-page="3"
            style="margin-top:20px;">
            <b-table-column label="ID" v-slot="props" width="50">
                {{props.row.id}}
            </b-table-column>
            <b-table-column label="Severity" v-slot="props" width="10">
                {{props.row.severity}}
            </b-table-column>
             <b-table-column label="Category" v-slot="props" width="10">
                <span>
                    <b-icon pack="fas" v-if="props.row.category == 'health'"
                        icon="heartbeat">
                    </b-icon>
                     <b-icon pack="mdi" v-if="props.row.category == 'dept'"
                        icon="cash">
                    </b-icon>
                     <b-icon pack="fas" v-if="props.row.category == 'nutrition'"
                        icon="food">
                    </b-icon>
                </span>
            </b-table-column>
            <b-table-column label="Description" v-slot="props">
                {{props.row.description}}
            </b-table-column>
            <b-table-column label="Date Modified" v-slot="props" width="200">
                {{props.row.date_modified}}
            </b-table-column>
            <b-table-column label="Edit" width="50" v-slot="props">
                 <b-button
                    size="is-small"
                    type="is-info"
                    icon-right="pencil"
                    @click="openCaseModal(props.row.id)" />
            </b-table-column>
        </b-table>

        <b-modal 
            v-model="isCaseModalActive"
            has-modal-card
            full-screen 
            :can-cancel="true">
            <case-modal v-bind="formProps"></case-modal>
        </b-modal>
    </div>
</template>

<script>
import CaseModal from '../components/CaseModal.vue'
export default {
    components: {
        CaseModal
    },
    methods:{
        openCaseModal : function(case_id){
            this.formProps.case_id = case_id
            this.isCaseModalActive = true
        }
    },
    data() {
        return {
            isCaseModalActive: false,
            formProps: {
                email: 'evan@you.com',
                password: 'testing',
                case_id: 'Undefined',
            },
            data: [
                { 'id': 1, 'category': 'health','description':'Reference site about Lorem Ipsum, giving information on its origins, as well as a random Lipsum generator.', 'severity': 0, 'date_modified': '2016-10-15 13:43:27'},
                { 'id': 2, 'category': 'health','description':'Reference site about Lorem Ipsum, giving information on its origins, as well as a random Lipsum generator.', 'severity': 1, 'date_modified': '2016-12-15 06:00:53'},
                { 'id': 3, 'category': 'dept','description':'Reference site about Lorem Ipsum, giving information on its origins, as well as a random Lipsum generator.', 'severity': 2, 'date_modified': '2016-04-26 06:26:28'},
                { 'id': 4, 'category': 'nutrition','description':'Reference site about Lorem Ipsum, giving information on its origins, as well as a random Lipsum generator.', 'severity': 1, 'date_modified': '2016-04-10 10:28:46'},
                { 'id': 5, 'category': 'nutrition','description':'Reference site about Lorem Ipsum, giving information on its origins, as well as a random Lipsum generator.', 'severity': 0, 'date_modified': '2016-12-06 14:38:38'}
            ],
        }
    },
}
</script>