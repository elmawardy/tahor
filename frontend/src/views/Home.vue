<template>
  <div class="home">
    <div class="columns" style="margin-top:50px;">
      <div class="column" style="text-align:end">
        <video-player :src="'http://techslides.com/demos/sample-videos/small.webm'"></video-player>
      </div>
      <div class="column">
        <div class="column" style="text-align:end">
          <section class="hero">
            <div class="hero-body">
              <div class="container">
                <h1 class="title has-text-primary">رسالة المؤسسة</h1>
                <h2
                  class="subtitle"
                >في بنك الشفاء المصري بنفتح حساب للأمل وكل يوم بنجدد مستشفى أو مستوصف رصيد الأمل فيه بيكبر. مفيش أصعب من الم مريض محتاج اتبرع الآن لبنك الشفاء المصري. تبرع. بنك الشفاء. تبرعات. اتبرع الآن. الأنواع: تبرعات, بنك الشفاء.</h2>
              </div>
            </div>
          </section>
        </div>
      </div>
    </div>
    <div class="columns">
      <div class="column has-text-centered" style="margin-top:50px;">
        <section class="hero">
          <div class="hero-body">
            <div class="container">
              <h1 class="title">شارك بمالك او وقتك</h1>
            </div>
          </div>
        </section>
      </div>
    </div>
    <div class="columns is-multiline">
      <div :key="index" v-for="(icase,index) in cases" class="column is-one-third">
        <CaseCard
          :org_name="icase.OrgName"
          :collected="icase.Collected"
          :target="icase.Target"
          :currency="icase.Currency"
          :comment="icase.Comment"
          :tags="icase.Tags"
          :updated_time="icase.DateModified"
          :case_id="icase.ID"
        ></CaseCard>
      </div>
    </div>

    <div class="has-text-centered" style="margin-top:30px;">
      <button class="button is-primary">مشاهدة المزيد</button>
    </div>
    <hr />
    <div class="columns">
      <div class="column has-text-centered" style="margin-top:10px;">
        <section class="hero">
          <div class="hero-body">
            <div class="container">
              <h1 class="title">الإنجازات</h1>
            </div>
          </div>
        </section>
      </div>
    </div>
    <div class="columns">
      <Accomplishment></Accomplishment>
      <Accomplishment></Accomplishment>
    </div>
    <div class="columns">
      <Accomplishment></Accomplishment>
      <Accomplishment></Accomplishment>
    </div>
    <div class="has-text-centered" style="margin-top:30px;">
      <button class="button is-primary">مشاهدة المزيد</button>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src
import VideoPlayer from "@/components/VideoPlayer.vue";
import CaseCard from "@/components/CaseCard.vue";
import Accomplishment from "@/components/Accomplishment.vue";

export default {
  name: "Home",
  components: {
    VideoPlayer,
    CaseCard,
    Accomplishment
  },
  data() {
    return {
      cases: []
    };
  },
  mounted() {
    fetch(this.$store.state.backendURL + "/api/case/select")
      .then(response => response.json())
      .then(response => {
        if (response.State == "Success") {
          this.cases = response.Cases;
        }
      });
  }
};
</script>
