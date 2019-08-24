<template>
  <div class="home">
    <WeatherDisplay />
    <h3>Angebote</h3>
    <InstagramImages hashtag="gornerbahn" />
    <h3>Aktivitäten</h3>
    <InstagramImages hashtag="mountainclimbing" />
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

<style lang="postcss"></style>
