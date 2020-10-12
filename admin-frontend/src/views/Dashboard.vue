<template>
<div style="padding:10px;min-height:100vh">
   <p class="title">Dashboard</p>
   <div class="columns">

      <div class="column">
         <div class="card">
            <div class="card-content">
               <div class="content">
                  <router-link to="/users">
                     <div class="columns">
                        <div class="column" style="text-align:center">
                              <b-icon
                                 pack="fas"
                                 icon="user"
                                 size="is-large"
                                 custom-size="fa-4x"
                                 type="is-primary">
                              </b-icon>
                        </div>
                           <div class="column" style="cursor:pointer;">
                                 <div class="columns is-gapless is-multiline">
                                    <div class="column is-three-quarters">
                                       <p class="title">Users</p>
                                    </div>
                                    <div class="column is-three-quarters">
                                       <p style="font-size:1rem">3</p>
                                    </div>
                                 </div>
                           </div>
                     </div>
                  </router-link>
               </div>
            </div>
         </div>
      </div>

       <div class="column" >
         <div class="card" style="position:relative;">
         <div style="position:absolute;top:0px;height:3px;background-color:green;width:100%;">
            &nbsp;
         </div>
            <div class="card-content">
               <div class="content">
                 <div class="columns">
                    <div class="column is-one-quarter" style="text-align:center">
                         <b-icon
                           pack="fas"
                           icon="heart"
                           size="is-large"
                           custom-size="fa-4x"
                           type="is-primary">
                        </b-icon>
                    </div>
                    <div class="column is-three-quarter">
                        <div class="columns is-gapless is-multiline">
                           <div class="column is-three-quarters">
                              <p class="title">Donations</p>
                           </div>
                           <div class="column">
                              <p style="font-size:1rem">500 <span style="font-size:10px">EGP</span>/100 <span style="font-size:10px">EGP</span>&nbsp;<span style="font-size:10px">(donations/needs)</span></p>
                           </div>
                        </div>
                    </div>
                 </div>
               </div>
            </div>
         </div>
      </div>

      <div class="column">
         <div class="card">
            <div class="card-content">
               <div class="content">
                  <div class="columns">
                    <div class="column" style="text-align:center">
                         <b-icon
                           pack="fas"
                           icon="heartbeat"
                           size="is-large"
                           custom-size="fa-4x"
                           type="is-primary">
                        </b-icon>
                    </div>
                    <div class="column">
                        <div class="columns is-multiline is-gapless">
                           <div class="column is-three-quarters">
                              <p class="title">Cases</p>
                           </div>
                           <div class="column is-three-quarters">
                              <p style="font-size:1rem">199</p>
                           </div>
                        </div>
                    </div>
                 </div>
               </div>
            </div>
         </div>
      </div>
   </div>   

   <div class="columns">
      <div class="column">
         <div class="card">
            <div class="card-content">
               <div class="media">
                  <div class="media-left">
                     <b-icon
                        pack="fas"
                        icon="chart-line"
                        size="is-large"
                        custom-size="fa-3x"
                        type="is-info">
                     </b-icon>
                  </div>
                  <div class="media-content">
                     <p class="title is-4">Statistics</p>
                  </div>
               </div>

               <div class="content">
                   <div class="columns">
                      <div class="column">
                         <VueApexCharts
                           type="donut"
                           width="100%"
                           :options="caseActiveClosedOptions"
                           :series="caseActiveClosedSeries"
                        ></VueApexCharts>
                      </div>
                      <div class="column">
                         <VueApexCharts
                           width="100%"
                           type="line"
                           :options="needsDonationStatsOptions"
                           :series="needsDonationStatsSeries"
                           ></VueApexCharts>
                      </div>
                      <div class="column">
                         <VueApexCharts
                           width="100%"
                           type="line"
                           :options="caseDonorStatsOptions"
                           :series="caseDonorStatsSeries"
                           ></VueApexCharts>
                      </div>
                   </div>
               </div>
            </div>
         </div>
      </div>
   </div>
</div>   
</template>

<script>
import VueApexCharts from "vue-apexcharts";
export default {
   components: {
    VueApexCharts,
  },
  data() {
     return {
        caseActiveClosedOptions: {
            labels: ["Active Cases", "Closed"],
            colors: ["#247BA0", "#98b796"],
            plotOptions: {
               pie: {
                  donut: {
                     size: "50"
                  }
               }
            }
         },
         caseActiveClosedSeries: [1,3],
         needsDonationStatsOptions:{
            chart: {
               height: 350,
               type: "line",
               stacked: false
            },
            dataLabels: {
               enabled: false
            },
            colors: ["#247BA0","#2A7825"],
            stroke: {
               width: [4, 4,4],
               curve: "smooth"
            },
            xaxis: {
               categories: [],
               labels: {
                  hideOverlappingLabels: true,
                  showDuplicates: false,
                  tickAmount: "dataPoints"
               }
            },
            yaxis: [
               {
                  opposite: true
               }
            ],
            tooltip: {
               shared: false,
               intersect: true,
               x: {
                  show: false
               }
            },
            legend: {
               horizontalAlign: "left",
               offsetX: 40
            }
         },
         needsDonationStatsSeries: [
            {
               name: "Needs",
               data: [1,2,3]
            },
             {
               name: "Donations",
               data: [4,2,3]
            },
         ],
         caseDonorStatsOptions: {
            chart: {
               height: 350,
               type: "line",
               stacked: false
            },
            dataLabels: {
               enabled: false
            },
            colors: ["#247BA0","#2A7825"],
            stroke: {
               width: [4, 4,4],
               curve: "smooth"
            },
            xaxis: {
               categories: [],
               labels: {
                  hideOverlappingLabels: true,
                  showDuplicates: false,
                  tickAmount: "dataPoints"
               }
            },
            yaxis: [
               {
                  opposite: true
               }
            ],
            tooltip: {
               shared: false,
               intersect: true,
               x: {
                  show: false
               }
            },
            legend: {
               horizontalAlign: "left",
               offsetX: 40
            }
         },
         caseDonorStatsSeries: [
            {
               name: "Active Cases",
               data: [20,25,30]
            },
             {
               name: "Donors",
               data: [10,20,15]
            },
         ],
     }
  },
}
</script>