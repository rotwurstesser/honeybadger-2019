<template>
  <div class="home p-5 pt-0 mb-6">
    <WeatherDisplay />
    Home!
  </div>
</template>

<script>
import WeatherDisplay from "../components/WeatherDisplay.vue";
export default {
  name: "home",
  components: { WeatherDisplay }
};
</script>

<style lang="postcss" scoped></style>
