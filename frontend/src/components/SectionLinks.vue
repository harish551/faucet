<template lang="pug">
.section
  .section-header
    h1 Community Links
  .section-main
    list-item(
      href="https://t.me/OmniFLixChat"
      target="_blank"
      title="Community Chat"
      subtitle="#OmniFlixChat")
    list-item(
      href="https://explorer.omniflix.network/"
      target="_blank"
      title="Testnet Explorer"
      subtitle="explorer.omniflix.network")
    list-item(
      href="https://omniflix.network"
      target="_blank"
      title="OmniFlix Website"
      subtitle="omniflix.network")
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
