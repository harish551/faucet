<template lang="pug">
.section
  .section-header
    h1 Community Links
  .section-main
    list-item(
      href="https://regen.network/"
      target="_blank"
      title="Website"
      subtitle="#regen.network")
    list-item(
      href="http://t.me/regennetwork_public"
      target="_blank"
      title="Telegram"
      subtitle="#regennetwork_public")
    list-item(
      href="https://github.com/regen-network"
      target="_blank"
      title="Source Code"
      subtitle="#regen-network")
    list-item(
      href="https://regen.vitwit.com/"
      target="_blank"
      title="Explorer"
      subtitle="#regen-network")
</template>

<script>
import ListItem from "../components/NiListItem.vue";
export default {
  name: "faucet",
  components: {
    ListItem
  },
  data: () => ({}),
  methods: {}
};
</script>
