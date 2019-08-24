<template>
  <div class="ig-image-slider" v-if="images">
    <img
      class="ig-image pl-2"
      v-for="(image, index) in images"
      :key="'igimage' + hashtag + index"
      :src="image"
      :alt="'instagram image of ' + hashtag"
    />
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "InstagramImages",
  data() {
    return {
      images: null
    };
  },
  props: {
    hashtag: {
      type: String,
      default: "gornergrat"
    }
  },
  created() {
    axios
      .get("http://5.9.112.47:1569/instagram?hashtag=" + this.hashtag)
      .then(result => {
        this.images = result.data;
      });
  }
};
</script>

<style lang="postcss" scoped>
.ig-image-slider {
  @apply flex overflow-scroll;
}
.ig-image {
  min-width: calc(100% - 1.5em);
  height: auto;
}
</style>
