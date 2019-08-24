<template>
  <div class="home">
    <WeatherDisplay />
    <h3>Angebote</h3>
    <InstagramImages hashtag="gornerbahn" />
    <h3>Aktivitäten</h3>
    <InstagramImages hashtag="mountainclimbing" />
    <h3>Interaktive Map</h3>
    <iframe
      class="p-5"
      src="https://openmaptiles.github.io/positron-gl-style/#8/46.0190/7.7460"
      width="100%"
      height="480"
    ></iframe>
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

<style lang="postcss">
.slider-toggler {
  @apply text-black;
}
</style>
