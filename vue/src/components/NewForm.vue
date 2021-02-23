<template>
	<div>
		<div class="container">
			<SpH3>New {{ type }}</SpH3>
            <SpInput
                v-model="pollTitle"
                type="text"
                placeholder="Title"
                :disabled="flight"
            />
            <div v-for="option in options" :key="option.id" >
                <sp-input placeholder="Option" v-model="option.title"/>
            </div>
            <SpButton
                :loading="flight"
                :disabled="!hasAddress || flight"
                @click="add"
            >
                Add Option
            </SpButton>
            <div class="spacer" />
            <SpButton
                :loading="flight"
                :disabled="!valid || !hasAddress || flight"
                @click="submit"
            >
                Create {{ type }}
            </SpButton>
			<SpTypeList :type="type" :path="path" :module="module" />
		</div>
	</div>
</template>

<style scoped>
@import '../styles/main.css';

.container {
	font-family: var(--sp-f-primary);
}
.button {
	display: inline-block;
}
.spacer{
    height: 10px;
}
</style>

<script>
import SpInput from './SpInput'
import SpH3 from './SpH3'
import SpButton from './SpButton'
import SpTypeList from './SpTypeList'
import { Type, Field } from 'protobufjs'
import { SigningStargateClient } from '@cosmjs/stargate'
import { Registry } from '@cosmjs/proto-signing'

export default {
	components: {
		SpInput,
		SpH3,
		SpButton,
		SpTypeList
	},
	props: {
		path: {
			type: String,
			default: ''
		},
		type: {
			type: String,
			default: ''
		},
		preflight: {
			type: Function,
			default: () => {
				return obj => obj
			}
		},
		module: {
			type: String,
			default: ''
		}
	},
	data: function() {
		return {
			flight: false,
            options: [],
            pollTitle: ""
		}
	},
	computed: {
		hasAddress() {
			return !!this.$store.state.cosmos.auth.account.address
		},
		valid() {
            let emptyStringArray = this.options.filter( o => o.title == "" );
            return (this.options.length > 1) && !(emptyStringArray.length) && this.pollTitle;
		}
	},
	methods: {
        add() {
            this.options = [...this.options, { title: "" }];
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
                MsgCreate.add(new Field("creator", 1, "string"));
                MsgCreate.add(new Field("title", 2, "string"));
                let optionsField = new Field("options", 3, "string");
                optionsField.repeated = true;
                MsgCreate.add(optionsField);
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
                        title: this.pollTitle,
                        options: this.options.map(o => o.title)
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
                this.options = []
                this.pollTitle = ""
			}
		}
	}
}
</script>
