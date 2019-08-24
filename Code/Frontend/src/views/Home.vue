<template>
  <div class="home p-5 pt-0 mb-6">
    <WeatherDisplay />
    <InstagramImages hashtag="zermatt" />
    Home!
  </div>
</template>

<script>
import InstagramImages from "../components/InstagramImages.vue";
import WeatherDisplay from "../components/WeatherDisplay.vue";
export default {
  name: "home",
  components: { WeatherDisplay, InstagramImages }
};
</script>

<style lang="postcss" scoped></style>
