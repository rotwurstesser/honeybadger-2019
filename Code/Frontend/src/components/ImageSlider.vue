<template>
  <div>
    <span
      class="text-right text-white block mb-2 mr-4 cursor-pointer slider-toggler z-10 relative"
      @click="showAll = !showAll"
    >
      {{
      showAll ? 'Weniger' : 'Mehr'
      }}
    </span>
    <div :class="showAll ? 'ig-image-slider-all' : 'ig-image-slider'" v-if="images">
      <div class="ig-image-container relative" v-for="(image, index) in images" :key="index">
        <img class="ig-image" :src="image.src" alt="image" />
        <span @click.prevent="changeFavorite">
          <svg-icon
            :name="svgIcon"
            class="absolute mr-6 mt-4 right-0 svg-icon svg-icon--icn-star-filled top-0 cursor-pointer"
          />
        </span>

        <span class="absolute bottom-0 px-4 text-white w-9/12 pr-0 text-lg">{{image.text}}</span>
      </div>
    </div>
  </div>
</template>

<script>
import SvgIcon from "@/components/SvgIcon";

export default {
  name: "ImageSlider",
  components: {
    SvgIcon
  },
  props: {
    images: {
      required: true,
      default: []
    }
  },
  data() {
    return {
      showAll: false,
      svgIcon: "icn-star-outline"
    };
  },
  methods: {
    changeFavorite() {
      if (this.svgIcon === "icn-star-outline") {
        this.svgIcon = "icn-star-filled";
      } else {
        this.svgIcon = "icn-star-outline";
      }
    }
  }
};
</script>

<style lang="postcss" scoped>
.ig-image-slider {
  @apply flex flex-row overflow-scroll;
}
.ig-image-slider-all {
  @apply flex flex-row flex-col;
  & .ig-image {
    min-width: 100%;
    max-height: 248px;
    @apply px-0;
    @apply mt-5;
  }
}
.ig-image {
  @apply px-2;
  min-width: calc(100% - 2em);
  height: 248px;
}
.ig-image-container {
  min-width: 248px;
  max-height: 248px;
}
</style>
