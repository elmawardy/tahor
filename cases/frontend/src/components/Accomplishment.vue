<template>
  <div class="column is-half">
    <div class="columns is-gapless">
      <div class="column is-half" style="position:relative;">
        <figure class="image is-square">
          <img
            src="https://cdn.dribbble.com/users/1778913/screenshots/6562748/dribbble-machucado1.jpg"
          />
        </figure>
        <p style="position:absolute;top:0;background-color:white;padding:10px;z-index:9">قبل</p>
        <div
          style="width:100%;height:100%;top:0;position:absolute;background-color:rgba(1,1,1,0.3);"
        ></div>
      </div>
      <div class="column is-half" style="position:relative;">
        <figure class="image is-square">
          <img src="https://miro.medium.com/max/750/1*xLKFtlXiTPif_hTAIlXRjw.jpeg" />
        </figure>
        <p style="position:absolute;top:0;background-color:white;padding:10px;">بعد</p>
      </div>
    </div>
    <div dir="rtl">
      <p>Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries.</p>
      <button style="margin:5px;" class="button is-twitter">
        <span class="icon" style="margin-left:5px;">
          <i class="fas fa-info"></i>
        </span>
        <span>المزيد من المعلومات</span>
      </button>
    </div>
  </div>
</template>