<template>
  <div v-if="images">
    <img
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

<style lang="scss" scoped></style>
