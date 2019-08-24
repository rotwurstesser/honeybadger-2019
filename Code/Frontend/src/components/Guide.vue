<template>
  <div class="bg-GornerGratGold h-screen">
    <div class="flex justify-end mr-6">
      <svg-icon class="mt-16" name="close" />
    </div>

    <weather-display />
    <div class="text-white mt-10 ml-6">
      <h1 class="font-bold text-3xl">Das Wetter wird sonniger!</h1>
      <span
        class="inline-block mt-4 text-2xl"
      >Perfekte Gelegenheit, um das Alpenpanorama zu geniessen!</span>
    </div>
    <div class="mt-20 opacity-50 text-2xl">
      <span>Schönwetterprogramm</span>
    </div>
  </div>
</template>

<script>
import WeatherDisplay from "@/components/WeatherDisplay";
import SvgIcon from "@/components/SvgIcon";

export default {
  name: "Guide",
  components: {
    WeatherDisplay,
    SvgIcon
  }
};
</script>

<style lang="postcss" scoped>
</style>