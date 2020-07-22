<template>
  <div>
    <div class="card">
      <div class="card-image">
        <figure class="image is-4by3">
          <img src="https://bulma.io/images/placeholders/1280x960.png" alt="Placeholder image" />
        </figure>
      </div>
      <div class="card-content">
        <div class="media">
          <div class="media-left">
            <figure class="image is-48x48">
              <img
                src="data:image/jpeg;base64,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"
                alt="Placeholder image"
              />
            </figure>
          </div>
          <div class="media-content">
            <p class="title is-4">Resala</p>
            <p class="subtitle is-6">@resalacharity</p>
          </div>
        </div>

        <div class="content">
          <progress
            class="progress is-medium is-primary show-value"
            value="20"
            max="100"
            text="21,500 / 100,000 EGP"
          ></progress>
          <p>عم فتحي محتاج علاج لمضر السكر</p>
          <a>@bulmaio</a>.
          <a href="#">#css</a>
          <a href="#">#responsive</a>
          <br />
          <time datetime="2016-1-1">11:09 PM - 1 Jan 2016</time>
          <div class="is-divider" data-content="شارك"></div>
          <div dir="rtl">
            <div class="buttons" style="text-align:end">
              <button class="button is-primary">
                <div style="padding-left:10px;">
                  <span class="icon">
                    <i class="fas fa-heart"></i>
                  </span>
                </div>
                <span>
                  <strong>تبرع</strong>
                </span>
              </button>
              <a href="/case/1" class="button">
                <span>قرائة المزيد</span>
                <div style="padding-right:7px;">
                  <span class="icon">
                    <i class="fas fa-info"></i>
                  </span>
                </div>
              </a>
              <div style="margin-left:5px;margin-right:5px;">
                <button class="button" data-tooltip="حفظ">
                  <span class="icon has-text-info">
                    <i class="fas fa-save"></i>
                  </span>
                </button>
              </div>
              <div style="margin-left:5px;">
                <button class="button" data-tooltip="نشر">
                  <span class="icon has-text-info">
                    <i class="fas fa-share-alt"></i>
                  </span>
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
progress.show-value {
  position: relative;
}

progress.show-value:after {
  content: attr(text);
  position: absolute;
  top: 0;
  left: 50%;
  -webkit-transform: translateX(-50%);
  -ms-transform: translateX(-50%);
  transform: translateX(-50%);
  font-size: calc(1rem / 1.5);
  line-height: 1rem;
}

progress.show-value.is-small:after {
  font-size: calc(0.75rem / 1.5);
  line-height: 0.75rem;
}

progress.show-value.is-medium:after {
  font-size: calc(1.25rem / 1.5);
  line-height: 1.25rem;
}

progress.show-value.is-large:after {
  font-size: calc(1.5rem / 1.5);
  line-height: 1.5rem;
}
</style>