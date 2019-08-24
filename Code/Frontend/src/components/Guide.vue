<template>
  <div class="guide h-full bg-GornerGratGold p-6 pb-20">
    <div class="flex justify-end mr-6">
      <router-link to="/">
        <svg-icon class="mt-16 text-white" name="close" />
      </router-link>
    </div>

    <weather-display :showTemperature="false" />
    <div class="text-white mt-8">
      <h1 class="font-bold text-3xl">Das Wetter wird sonniger!</h1>
      <span class="inline-block mt-4 text-2xl">
        Perfekte Gelegenheit, um das Alpenpanorama zu geniessen!
      </span>
    </div>
    <div class="mt-20 opacity-50 text-2xl negative-margin">
      <span class="font-bold">Schönwetterprogramm</span>
    </div>

    <div class="guide-slider -m-2">
      <router-link to="/detail">
        <image-slider :images="images" />
      </router-link>
    </div>
  </div>
</template>

<script>
import WeatherDisplay from "@/components/WeatherDisplay";
import SvgIcon from "@/components/SvgIcon";
import ImageSlider from "@/components/ImageSlider";
import Picture1 from "@/assets/bitmap.png";

export default {
  name: "Guide",
  components: {
    WeatherDisplay,
    SvgIcon,
    ImageSlider
  },
  data() {
    return {
      images: [Picture1]
    };
  }
};
</script>

<style lang="postcss">
.guide {
  & .negative-margin {
    margin-bottom: -1.75rem;
  }

  & .ig-image-slider-all {
    @apply mt-4;
  }

  & img.ig-image {
    min-width: inherit;
    width: 248px;
    height: auto;
    @apply rounded-squarecard;
  }
}
</style>
