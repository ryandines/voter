<template>
  <div>
    <sp-input placeholder="Title" v-model="pollTitle" />
    <div v-for="option in options" :key="option.id" >
      <sp-input placeholder="Option" v-model="option.title"/>
    </div>
    <sp-button @click.native="add">+ Add option</sp-button>
    <div class="spacer" />
    <sp-button @click.native="submit">Create poll</sp-button>
  </div>
</template>
<style scoped>
    .spacer{
        height: 10px;
    }
</style>
<script>
import * as sp from "@tendermint/vue";
export default {
  components: { ...sp },
  data() {
    return {
      pollTitle: "",
      options: []
    };
  },
  methods: {
    add() {
      this.options = [...this.options, { title: "" }];
    },
		title(string) {
			return string.charAt(0).toUpperCase() + string.slice(1)
		},
		async submit() {
			if (this.valid && !this.flight && this.hasAddress) {
				const { RPC } = this.$store.state.cosmos.env.env
				const wallet = this.$store.getters['cosmos/wallet']
				const account = this.$store.getters['cosmos/account']
				const from_address = account.address
				const type = this.type.charAt(0).toUpperCase() + this.type.slice(1)
				const typeUrl = `/${this.path}.MsgCreate${type}`
				let MsgCreate = new Type(`MsgCreate${type}`)
				this.fields.forEach(f => {
					MsgCreate = MsgCreate.add(new Field(f[0], f[1], f[2]))
				})
				const registry = new Registry([[typeUrl, MsgCreate]])
				const client = await SigningStargateClient.connectWithWallet(
					RPC,
					wallet,
					{ registry }
				)
				const msg = {
					typeUrl,
					value: {
						creator: from_address,
						...this.fieldsList
					}
				}
				const fee = {
					amount: [{ amount: '0', denom: 'token' }],
					gas: '200000'
				}
				this.flight = true
				try {
					const path = this.path.replace(/\./g, '/')
					await client.signAndBroadcast(from_address, [msg], fee)
					this.$store.dispatch('cosmos/entityFetch', {
						type: this.type,
						path: path
					})
				} catch (e) {
					console.log(e)
				}
				this.flight = false
				Object.keys(this.fieldsList).forEach(f => {
					this.fieldsList[f] = ''
				})
			}
		}
  }
};

</script>