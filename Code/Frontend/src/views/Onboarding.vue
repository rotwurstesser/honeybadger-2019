<template>
  <div>
    <div class="onboarding p-6">
      <div class="steps mx-16 mb-6">
        <div @click="step = 1" :class="step === 1 ? 'step--active' : ''" class="step"></div>
        <div @click="step = 2" :class="step === 2 ? 'step--active' : ''" class="step"></div>
        <div @click="step = 3" :class="step === 3 ? 'step--active' : ''" class="step"></div>
      </div>
      <div v-if="step === 1" class="questionWrap">
        <h2 class="text-GornerGratGold text-xl text-center font-bold mb-2">Was trifft zu</h2>
        <h3 class="text-GornerGratGold text-lg text-center">Ich bin</h3>
        <div class="choice-row">
          <div @click.prevent="setNextStep(2)" class="choice">
            <svg-icon class="text-black" name="tourist" />
            <span class="choice-text">Tourist</span>
          </div>
          <div @click.prevent="setNextStep(2)" class="choice">
            <svg-icon class="text-black" name="swiss" />
            <span class="choice-text">Schweizer</span>
          </div>
        </div>
      </div>
      <div v-if="step === 2" class="questionWrap">
        <h2 class="text-GornerGratGold text-xl text-center font-bold mb-2">Mit wem bist du hier?</h2>
        <h3 class="text-GornerGratGold text-lg text-center">Wähle aus</h3>
        <div class="choice-row">
          <div @click.prevent="setNextStep(3)" class="choice">
            <svg-icon class="text-black" name="trekking-person" />
            <span class="choice-text">Alleine</span>
          </div>
          <div @click.prevent="setNextStep(3)" class="choice">
            <svg-icon class="text-black" name="family-child" />
            <span class="choice-text">Mit der Familie</span>
          </div>
        </div>
        <div class="choice-row">
          <div @click.prevent="setNextStep(3)" class="choice">
            <svg-icon class="text-black" name="family-home" />
            <span class="choice-text">Mit Freunden</span>
          </div>
          <div @click.prevent="setNextStep(3)" class="choice">
            <svg-icon class="text-black" name="dating-couple-balloon" />
            <span class="choice-text">Mit dem Partner</span>
          </div>
        </div>
      </div>
      <div v-if="step === 3" class="questionWrap">
        <h2
          class="text-GornerGratGold text-xl text-center font-bold mb-2"
        >Das würde ich gerne machen</h2>
        <h3 class="text-GornerGratGold text-lg text-center">Wähle aus</h3>
        <div class="choice-row">
          <div @click.prevent="activate()" class="choice">
            <svg-icon class="text-black" name="trekking-person" />
            <span class="choice-text">Wandern</span>
          </div>
          <div @click.prevent="activate()" class="choice">
            <svg-icon class="text-black" name="family-child" />
            <span class="choice-text">Fotografieren</span>
          </div>
        </div>
        <div class="choice-row">
          <div @click.prevent="activate()" class="choice">
            <svg-icon class="text-black" name="family-home" />
            <span class="choice-text">Klettern</span>
          </div>
          <div @click.prevent="activate()" class="choice">
            <svg-icon class="text-black" name="dating-couple-balloon" />
            <span class="choice-text">Essen</span>
          </div>
        </div>
      </div>
    </div>
    <footer class="absolute bottom-0 w-full">
      <svg-icon :name="'mountainfooter' + step" class="w-full h-auto" />
      <span
        class="absolute bottom-0 mb-10 ml-10 text-white text-2xl"
        @click.prevent="goToNextStep()"
      >Skip</span>
      <span
        v-if="step === 3"
        class="absolute bottom-0 right-0 mr-10 mb-10 text-white text-2xl"
        @click.prevent="goToHome()"
      >Weiter -></span>
    </footer>
  </div>
</template>

<script>
import SvgIcon from "../components/SvgIcon.vue";
export default {
  components: {
    SvgIcon
  },
  methods: {
    setNextStep(step) {
      const event2 = event.currentTarget;
      event.currentTarget.classList.toggle("choice--active");
      // give the user some time to see the coloring/let the eyes settle
      setTimeout(() => {
        event2.classList.toggle("choice--active");
        this.step = step;
      }, 100);
    },
    activate() {
      event.currentTarget.classList.toggle("choice--active");
    },
    goToNextStep() {
      if (this.step === 3) {
        this.goToHome();
        return;
      }
      this.step++;
    },
    goToHome() {
      this.$router.push("home");
    }
  },
  data() {
    return {
      step: 1
    };
  }
};
</script>

<style lang="postcss">
.onboarding {
  & h3 {
    @apply mb-16;
  }
  & .wrapping {
    top: 50%;
    left: 50%;
    transform: translateY(-50%, -50%);
  }
  & .choice-row {
    @apply flex flex-row justify-around p-2 mx-2;
    & .choice {
      width: 150px;
      height: 150px;
      margin: 0 auto;

      @apply mx-1 shadow-md rounded-squarecard bg-white text-center flex flex-col justify-center;
      & svg {
        @apply stroke-black;
      }

      & .choice-text {
        @apply font-semibold text-lg mt-2;
      }
    }
    & .choice--active {
      @apply bg-GornerGratGold;
      & span,
      & svg {
        @apply text-white stroke-white;
      }
    }
    & svg {
      width: 40px;
      height: 40px;
      margin: 0 auto;
    }
  }
  & .steps {
    @apply flex flex-row justify-between;
  }

  & .step {
    width: 30%;
    height: 5px;
    @apply bg-GornerGratGrey3;
    &--active {
      @apply bg-GornerGratGold;
    }
  }
  & .step--active {
    @apply bg-GornerGratGold;
  }
}
footer {
  height: 140px;
}
</style>
