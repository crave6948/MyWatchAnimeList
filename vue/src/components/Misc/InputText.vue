<template>
    <div class="input-group">
        <label class="input-group__label">{{ getText }}</label>
        <input v-model="inputText" type="text" class="input-group__input" @input="emitInputText"/>
    </div>
</template>

<script>
export default {
    name: 'InputText',
    props: {
        msg: String
    },
    data() {
        return {
            inputText: ''
        }
    },
    computed: {
        getText() {
            return this.$props.msg
        }
    },
    methods: {
        emitInputText() {
            this.$emit('input-text', this.inputText);
        }
    }
}
</script>

<style scoped>
* {
    box-sizing: border-box;
}

@media (prefers-reduced-motion: reduce) {
    * {
        animation: none !important;
        transition-duration: 0.001s !important;
    }
}

@keyframes background-move {
    0% {
        background-position: 0 0, 0 0;
    }

    100% {
        background-position: 100vw 0, 0 0;
        background-position: max(100vw, 40em) 0, 0 0;
    }
}

@supports (mix-blend-mode: darken) {
  .input-group {
    position: relative;
  }
  .input-group__label {
    position: absolute;
    left: 3em;
    top: -0.28em;
    background: #000;
  }
}

.input-group {
    width: 100%;
    max-width: 20em;
    display: flex;
    flex-direction: column;
    z-index: 1;
}
.input-group__label {
  padding: 0 0.5em;
  margin-bottom: 0.5em;
  text-transform: uppercase;
  font-size: 0.875em;
  letter-spacing: 0.1em;
  color: #ccd;
  color: rgba(255, 220, 255, 0.6);
  cursor: pointer;
}
.input-group__input {
    color: #fff;
    font-size: 1.25rem;
    line-height: 1;
    border-style: none;
    outline: none;
    height: calc(1em + 1.6em + 0.5em);
    width: 100%;
    padding: 0.8em 1em;
    border: 0.25em solid transparent;
    background-image: linear-gradient(#000, #000), linear-gradient(120deg, #f09 0%, #0ff 50%, #9f0 100%);
    background-origin: border-box;
    background-clip: padding-box, border-box;
    border-radius: 1.8em;
    background-size: 200% 100%;
    transition: background-position 0.8s ease-out;
}

.input-group__input:hover {
    background-position: 100% 0;
}</style>