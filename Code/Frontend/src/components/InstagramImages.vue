<template>
  <div>
    <image-slider :images="images" />
  </div>
</template>

<script>
import axios from "axios";
import ImageSlider from "@/components/ImageSlider";
export default {
  name: "InstagramImages",
  components: {
    ImageSlider
  },
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
        this.images = result.data.map(element => {
          return { src: element, text: "" };
        });
      });
  }
};
</script>

<style lang="postcss" scoped>
.ig-image-slider {
  @apply flex flex-row overflow-scroll -m-2;
}
.ig-image-slider--all {
  @apply flex flex-row flex-col;
  .ig-image {
    min-width: 100%;
    @apply px-0;
  }
}
.ig-image {
  @apply p-2;
  min-width: calc(100% - 2em);
  height: auto;
}
</style>
