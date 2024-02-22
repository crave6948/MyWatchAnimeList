<template>
  <div id="stickyDiv">
    <img @click="clicked" id="icon" src="../../../assets/plusicon.svg" alt="plus" />
  </div>
  <div id="form" :class="[{disabled: !open}]">
    <div id="panel" :class="[{ in: open }, { out: !open }]">
      <ActionForm :sent="calladd" />
    </div>
  </div>
</template>

<script>
var lastMS = 0
import ActionForm from './ActionForm.vue'
export default {
  props: ['addAnime'],
  components: {
    ActionForm
  },
  data() {
    return {
      open: false,
      allowedClick: false
    }
  },
  methods: {
    calladd(animeName, animeWatch, animeImage) {
      this.addAnime(animeName, animeWatch, animeImage)
    },
    clicked() {
      if (Date.now() - lastMS < 500) return
      this.open = !this.open
      lastMS = Date.now()
    }
  }
}
</script>

<style scoped>
#icon {
  width: 50px;
  height: 50px;
  border-radius: 12px;
  position: relative;
}

#stickyDiv {
  position: fixed;
  bottom: 5%;
  right: 5%;
  z-index: 1;
}

#buttonContainer {
  display: flex;
  justify-content: center;
  margin-top: 20px;
}

#form {
  position: fixed;
  left: 50%;
  top: 50%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  transform: translate(-50%, -50%);
}

#panel {
  border-radius: 12px;
  box-shadow: 0px 0px 10px 0px rgba(0, 0, 0, 0.5);
  backdrop-filter: blur(10px);
}

#form.disabled {
  display: none;
}

.in {
  animation: fadeIn 0.2s forwards;
}
@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateX(-50px);
  }
  to {
    opacity: 1;
  }
}
.out {
  animation: fadeOut 0.2s forwards;
}
@keyframes fadeOut {
  from {
    opacity: 1;
  }
  to {
    opacity: 0;
    transform: translateX(-50px);
  }
}
</style>
