<template>
  <div class="detail">
    <img class="opacity-75" :src="require('@/assets/detailasset.png')" alt="" />
    <router-link to="/Guide">
      <svg-icon
        class="absolute mr-12 mt-24 right-0 svg-icon text-white top-0"
        name="close"
      />
    </router-link>
    <div class="-mt-56 relative videobox mx-6">
      <div
        v-if="paused"
        class="darkener rounded-squarecard absolute w-full h-full"
      ></div>
      <video
        class="rounded-squarecard"
        id="headervid"
        :src="require('@/assets/offer.mp4')"
      ></video>
      <span class="play-pause-wrap" @click.prevent="playPause">
        <svg-icon class="videobutton" :name="paused ? 'play' : 'pause'" />
      </span>
    </div>

    <div class="p-6">
      <h3>
        Riffelsee, der Seelen-Spiegel des Matterhorns oberhalb von Zermatt
      </h3>
      <div class="actionbutton ">
        <svg-icon class="mr-2" name="camera" /><span>
          Fotopoint:Riffelsee Spiegelung
        </span>
      </div>
      <div class="actionbutton ">
        <svg-icon class="mr-2" name="train" />
        <span>Anreise: 10min ab Station Rotenboden</span>
      </div>
      <div class="actionbutton ">
        <svg-icon class="mr-2" name="flower" />
        <span>Seltene Blumen und Pflanzen</span>
      </div>
      <div class="actionbutton ">
        <svg-icon class="mr-2" name="climbing-mountain" />
        <span>Ausgangspunkt für Kletterrouten am Riffelhorn</span>
      </div>
      <p>
        Mit der Gornergrat Bahn fahren Sie direkt ab Zermatt bis zur Station
        Rotenboden, die vorletzte Station vor dem Gornergrat. Ab Bahnhof
        Rotenboden nehmen Sie den Wanderweg und marschieren auf dem Riffelseeweg
        10min bis Sie den einmaligen See erreichen. Die kurze Wanderung bergab
        ist einfach zu meistern und braucht keine Wanderkenntnisse oder
        spezielle Wanderausrüstung. Bitte bleiben Sie auf den gekennzeichneten
        Wanderrouten, um die Pflanzen- und Tierwelt des Naturschutzgebietes zu
        schützen. Zurück nach Zermatt bringt Sie die Gornergrat Bahn.
      </p>
    </div>
  </div>
</template>

<script>
import SvgIcon from "../components/SvgIcon.vue";
export default {
  components: {
    SvgIcon
  },
  name: "Detail",
  data() {
    return {
      paused: true
    };
  },
  methods: {
    playPause() {
      this.paused = !this.paused;
      const vid = document.querySelector("#headervid");
      if (this.paused) {
        vid.pause();
      } else {
        vid.play();
      }
    }
  }
};
</script>

<style lang="postcss">
.detail {
  & .actionbutton {
    @apply p-5 py-6 w-full bg-GornerGratGold text-white flex flex-row my-5 font-bold;
    & svg {
      @apply self-center;
    }
    border-radius: 5px;
    box-shadow: 0 0 12px 0 rgba(0, 0, 0, 0.06);
  }
  & h3 {
    @apply text-3xl font-bold;
  }
  & .darkener {
    background: rgba(133, 124, 124, 0.5);
  }
  & .play-pause-wrap {
    @apply cursor-pointer absolute z-10;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
  }
  & .videobutton {
    @apply text-white;
    width: 55px;
    height: 55px;
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
