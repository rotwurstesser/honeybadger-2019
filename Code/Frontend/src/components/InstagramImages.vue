<template>
  <div>
    <span
      class="text-right text-white block mb-2 mr-4 cursor-pointer slider-toggler"
      @click="showAll = !showAll"
    >
      {{ showAll ? "Weniger Anzeigen" : "Mehr Anzeigen" }}
    </span>
    <div
      :class="showAll ? 'ig-image-slider-all' : 'ig-image-slider'"
      v-if="images"
    >
      <img
        class="ig-image"
        v-for="(image, index) in images"
        :key="'igimage' + hashtag + index"
        :src="image"
        :alt="'instagram image of ' + hashtag"
      />
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  name: "InstagramImages",
  data() {
    return {
      images: null,
      showAll: false
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
  @apply flex flex-row overflow-scroll;
}
.ig-image-slider--all {
  @apply flex flex-row flex-col;
  .ig-image {
    min-width: 100%;
    @apply px-0;
  }
}
.ig-image {
  @apply px-2;
  min-width: calc(100% - 1.5em);
  height: auto;
}
</style>
