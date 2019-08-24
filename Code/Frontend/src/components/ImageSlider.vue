<template>
  <div>
    <span
      class="text-right text-white block mb-2 mr-4 cursor-pointer slider-toggler"
      @click="showAll = !showAll"
    >{{ showAll ? "Weniger" : "Mehr" }}</span>
    <div :class="showAll ? 'ig-image-slider-all' : 'ig-image-slider'" v-if="images">
      <img
        class="ig-image"
        v-for="(image, index) in images"
        :key="'igimage' + index"
        :src="image"
        :alt="'image'"
      />
    </div>
  </div>
</template>

<script>
export default {
  name: "ImageSlider",
  props: {
    images: {
      required: true,
      default: []
    },
    showAll: {
      required: true,
      default: true
    }
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
