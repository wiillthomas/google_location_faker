
<script>
	import { fade } from "svelte/transition"
	
	let location;
	let query;
	let googleLink;
	let flashError;
	let locationPlaceholder = "Fr";
	let possibleLocations = []

	function handleClick(){
		if ( location && query ) {
			fetch(`http://localhost:8080/api?q=${ query }&location=${ location }` )
			.then( ( response ) =>  response.json() )
			.then( ( data ) => {
				googleLink = `https://google.com/search?${data}`
				flashError = null;
			} )
			.catch( err => {
				flashError = "An issue occoured while communicating with the API. Email help@thisdomain.com if the problem persists."
			} )
		} else {
			flashError = "Please fill in the required fields"
		}
	}

	function handleAutoComplete() {
		console.log(location)
		fetch(`http://localhost:8080/autocomplete?input=${location}`)
		.then(( response ) => {
			return response.json()
		})
		.then((data) => {
			locationPlaceholder = ""
			possibleLocations.length = 0
			for ( let i = 0; i < 10; i += 1 ) {
				if ( data[i] ){
					possibleLocations.push(`${location}${data[i]}`)
				}
			}
		})
		.catch(err =>  console.log(err))
	}

</script>

<style>
	.flash-message {
		display: flex;
		height: 50px;
		background: #bb0000;
		color: white;
		position: absolute;
		top: 0px;
		opacity: 1;
		left: 0px;
		width: 100vw;
		justify-content: center;
		align-items: center;
	}

	.container {
		height: 100vh;
		width: 100vw;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
	}
	
	.container .input-container {
		display: flex;
		flex-direction: row;
	}
	.container input, .container button {
		height: 40px;
	}

	input.query {
		position: relative;
	}
	input.query::before {
		content: "Query";
		position: absolute;
		top: -10px;
		left: 0;
	}
	
	.possible-location-container {
		position: absolute;
		display: flex;
		flex-direction: column
	}

	.possible-location-container button {
		text-align: left;
	}

</style>

{#if flashError}
	<div class="flash-message">
		{ flashError }
	</div>
{/if}

<div class="container">
	<div class="input-container">
		<input class="query" bind:value={query}>
		<div class="location location-input__container">
			<input bind:value={location} on:keyup={handleAutoComplete} >
				<div class="possible-location-container">
					{#each possibleLocations as location}
						<button name={location}>
							{location}
						</button>
					{/each}
				</div>
		</div>
		<button on:click={handleClick}>
			Get Results
		</button>
	</div>
	<div class="link-container">
		{#if googleLink}
			<a href={googleLink}>link here!</a>
		{/if}
	</div>
</div>