<template lang="pug">
#faucet
  .section
    faucet-header
    form(v-on:submit.prevent='onSubmit', method='post')
      form-group(:error='$v.fields.response.$error'
        field-id='faucet-response' field-label='Captcha')
        vue-recaptcha#faucet-response(
          ref="recaptcha"
          @verify="onVerify"
          @expired="onExpired"
          :sitekey="config.recaptchaSiteKey")
        form-msg(name='Captcha' type='required' v-if='!$v.fields.response.required')
      form-group(:error='$v.fields.selectedChain.$error' field-id='faucet-chain' field-label='Select Chain')
        select.custom-select(v-model='fields.selectedChain' :options="options")
          option(value="") Select Chain
          option(v-for="option in options" v-bind:value="option.value") {{option.text}}
        form-msg(name='Chain ' type='required' v-if='!$v.fields.selectedChain.required')
      
      form-group(:error='$v.fields.address.$error'
        field-id='faucet-address' field-label='Send To')
        field#faucet-address(
          type='text'
          v-model='fields.address'
          placeholder='Your account address'
          size="lg")
        form-msg(name='Address' type='required' v-if='!$v.fields.address.required')
        form-msg(name='Address' type='bech32' :body="bech32error" v-else-if='!$v.fields.address.bech32Validate')
      form-group
        btn(v-if='sending' value='Sending...' disabled color="primary" size="lg")
        btn(v-else @click='onSubmit' value="Send me tokens" color="primary" size="lg" icon="send")
  section-links
</template>

<script>
import axios from "axios";
import VueRecaptcha from "vue-recaptcha";
import { mapGetters } from "vuex";
import { required } from "vuelidate/lib/validators";
import Btn from "@nylira/vue-button";
import Field from "@nylira/vue-field";
import FormGroup from "../components/NiFormGroup";
import FormMsg from "../components/NiFormMsg";
import FaucetHeader from "../components/FaucetHeader";
// import SectionJoin from "../components/SectionJoin.vue";
import SectionLinks from "../components/SectionLinks.vue";
export default {
  name: "faucet",
  components: {
    Btn,
    Field,
    FormGroup,
    FaucetHeader,
    FormMsg,
    SectionLinks,
    VueRecaptcha
  },
  computed: {
    ...mapGetters(["config"])
  },
  data: () => ({
    fields: {
      response: "",
      address: "",
      selectedChain: ""
    },
    selectedChain: "",
    options: [
      { text: "testnetibc0 (stake)", value: "testnetibc0" },
      { text: "testnetibc1 (stake)", value: "testnetibc1" },
      { text: "vitwit (stake)", value: "vitwit" },
      { text: "vitwitibc2 (stake)", value: "vitwitibc2" },
      { text: "dokia-bombers (stake)", value: "dokia-bombers" },
      { text: "chainl1 (stake)", value: "chainl1" },
      { text: "ibc-corestart (stake)", value: "ibc-corestart" }
    ],
    sending: false
  }),
  methods: {
    resetForm() {
      this.fields.address = "";
      this.fields.response = "";
      this.fields.selectedChain = "";
      this.$refs.recaptcha.reset();
      this.$v.$reset();
    },
    onVerify(response) {
      this.fields.response = response;
    },
    onExpired() {
      this.$refs.recaptcha.reset();
    },
    async onSubmit() {
      this.$v.$touch();
      if (this.$v.$error) return;

      this.sending = true;

      var data = {
        address: this.fields.address.toString(),
        response: this.fields.response.toString(),
        selectedChain: this.fields.selectedChain.toString()
      };

      var bodyFormData = new FormData();
      bodyFormData.set("address", data.address);
      bodyFormData.set("response", data.response);
      bodyFormData.set("chain", data.selectedChain);

      axios(
        // .post(this.config.claimUrl, data)
        {
          method: "post",
          url: this.config.claimUrl,
          data: bodyFormData,
          config: { headers: { "Content-Type": "multipart/form-data" } }
        }
      )
        .then(() => {
          this.$refs.recaptcha.reset();
          this.sending = false;
          this.$store.commit("notify", {
            title: "Successfully Sent",
            body: `Sent tokens to ${this.fields.address}`
          });
          this.resetForm();
        })
        .catch(err => {
          this.$refs.recaptcha.reset();
          var msg = err.message;
          if (err.response && err.response.data && err.response.data.message) {
            msg = err.response.data.message;
          }

          this.sending = false;
          this.$store.commit("notifyError", {
            title: "Error Sending",
            body: `An error occurred while trying to send: "${msg}"`
          });
        });
    },
    bech32Validate(param) {
      try {
        if (param.length == 45 && param.startsWith("emoney")) {
          this.bech32error = null;
          return true;
        } else {
          this.bech32error = "Invalid address";
          return false;
        }
      } catch (error) {
        this.bech32error = error.message;
        return false;
      }
    }
  },
  validations() {
    return {
      fields: {
        address: {
          required,
          bech32Validate: this.bech32Validate
        },
        selectedChain: { required },
        response: { required }
      }
    };
  }
};
</script>

<style lang="stylus">
@import '~variables';

#faucet {
  max-width: 60rem;
  width: 100%;
  margin: 0 auto;
}

.custom-select {
    position: relative;
    font-family: Arial;
    min-width: 300px;
    /* min-height: 50px; */
    border-radius: 2px !important;
    height: 3rem;
    font-size: 1.125rem;
    padding-left: 0.75rem;
    padding-right: 0.75rem;
}

.section {
  margin: 0.5rem;
  padding: 1rem;
  background: var(--app-bg);
  position: relative;
  z-index: 10;

  label {
    display: none;
  }

  input:-webkit-autofill {
    -webkit-text-fill-color: var(--txt) !important;
    -webkit-box-shadow: 0 0 0px 3rem var(--app-fg) inset;
  }

  .section-main {
    padding: 0 1rem;
  }
}

@media screen and (min-width: 375px) {
  .section {
    padding: 2rem 1rem;
  }
}

@media screen and (min-width: 768px) {
  .section {
    padding: 3rem 2rem;
    margin: 1rem;
  }
}
</style>
