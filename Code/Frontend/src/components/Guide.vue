<template>
  <div class="bg-GornerGratGold">
    <div class="flex justify-end mr-6">
      <svg-icon name="close" />
    </div>

    <weather-display />

    <h1 class="font-bold text-2xl ml-6 text-white">Das Wetter wird sonniger!</h1>
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