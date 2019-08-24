<template>
  <div class="home m-6 ">
    <WeatherDisplay />
    <h3 class="-mb-8">Angebote</h3>
    <InstagramImages hashtag="gornerbahn" />
    <h3 class="-mb-8 mt-12">Aktivitäten</h3>
    <InstagramImages hashtag="mountainclimbing" />
    <h3 class="mt-12">Interaktive Map</h3>
    <iframe
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
.home {
  & h3 {
    @apply text-3xl font-bold;
  }
  & .weather-component {
    & span,
    & svg {
      @apply text-black;
    }
    & span {
      font-size: 35px;
      @apply ml-2;
    }
    & .weathersvgs {
      width: 60px;
      margin-top: 0;
    }
  }
  & .slider-toggler {
    color: black;
  }
}
</style>
