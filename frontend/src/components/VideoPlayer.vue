<template>
  <div>
    <video controls>
      <source :src="src" type="video/webm" />
    </video>
  </div>
</template>

<script>
export default {
  props: ["src"]
};
</script>