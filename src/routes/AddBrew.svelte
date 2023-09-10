<script lang="ts">
	import { tobrews } from '$lib/tobrews';
	import type { ToBrew, Brew } from '$lib/types';

	const addBrew = async () => {
		const response = await fetch('http://localhost:3333/tobrews/', {
			method: 'POST',
			body: JSON.stringify({
				name: 'new brew',
				roaster: { String: 'S&W', Valid: true },
				link: { String: 'https://us.mystery.coffee', Valid: true },
				brewed: false,
				timeToBrew: new Date().toISOString()
			}),
			headers: {
				'Content-Type': 'application/json',
				Origin: 'http://localhost:5173/'
			}
		});

		const brew = (await response.json()) as Brew;
		console.log(brew);

		let newBrew: ToBrew = {
			id: brew.id,
			name: 'new brew',
			bean: '',
			brewed: false,
			time: new Date(),
			data: {}
		};

		tobrews.update((brews) => [newBrew, ...brews]);
	};
</script>

<button on:click={addBrew}>+ Add Brew</button>
