
<script>
	let location;
	let query;
	let googleLink;
	let flashError;
	let locationPlaceholder = "Fr";
	let possibleLocations = []
	let possibleLocationFocus = -1;
	let autoCompleteError = false;
	let displayPossibleLocations = false;

	function handleClick(){
		displayPossibleLocations = false;
		if ( location && query && location.length > 3 ) {
			fetch(`https://location-faker.onrender.com/api?q=${ query }&location=${ location }` )
			.then( ( response ) =>  response.json() )
			.then( ( data ) => {
				googleLink = `https://google.com/search?${data}`
				flashError = null;
			} )
			.catch( err => {
				flashError = "An issue occurred while communicating with the API."
			} )
		} else if ( location.length < 4 ) {
			flashError = "Location must be more than 3 characters"
		} else {
			flashError = "Please fill in the required fields"
		}
	}

	function handlePossibleLocationSubmit( e ) {
		e.preventDefault() 
		location = e.currentTarget.name
		displayPossibleLocations = false;
	}

	function handleLocationInputChange( e ) {
		googleLink = ""
		displayPossibleLocations = true

		if ( e.keyCode == 38 ) {
			if ( possibleLocationFocus !== -1 ){
				e.preventDefault()
				possibleLocationFocus -= 1
				document.getElementById(`possibleLocation--${ possibleLocationFocus }`).focus()
			}
		} else if ( e.keyCode == 40 ) {
			e.preventDefault()
			possibleLocationFocus += 1
			document.getElementById(`possibleLocation--${ possibleLocationFocus }`).focus()
		} else {
			possibleLocations.length = 0
			possibleLocationFocus = -1
			fetch(`https://location-faker.onrender.com/autocomplete?input=${location}`)
			.then(( response ) => {
				autoCompleteError = false
				return response.json()
			})
			.then((data) => {
				locationPlaceholder = ""
				possibleLocations.length = 0
				for ( let i = 0; i < 4; i += 1 ) {
					if ( data[i] ){
						possibleLocations.push(`${location}${data[i]}`)
					}
				}

			})
			.catch(( err ) =>  {
				displayPossibleLocations = false
			})
		}
		
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
		min-height: 100vh;
		min-width: 100vw;
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;

		background: linear-gradient(270deg, #e74c3c, #c0392b);
		background-size: 400% 400%;

		-webkit-animation: Background-Gradient 30s ease infinite;
		-moz-animation: Background-Gradient 30s ease infinite;
		animation: Background-Gradient 30s ease infinite;
	}
	
	.container .input-container {
		display: flex;
		flex-direction: row;
		flex-wrap: wrap;
		padding: 10px;
	}

	.container .input-container > * {
		margin: 0 0 25px;
		width: 250px;
	}
	.container input {
		padding: 15px 15px 15px 15px;
		color: white;
		background: transparent;
		border: 1px solid white;
		cursor: pointer;
		transition: 0.3s;
		width: 250px;
		text-align: left;
	}
	
	.container button {
		padding: 15px 15px 25px 15px;
		background: white;
		color: #c0392b;
		border: 1px solid white;
		cursor: pointer;
		transition: 0.3s;
		text-align: left;
		top: 5px;
		position: relative;
		width: 150px;
		height: 49px;
		border-radius: 0;
		font-size: 14px;
	}
	

	input.query {
		position: relative;
	}

	.location-input__container:focus-within .possible-location-container {
		position: absolute;
		display: flex;
		flex-direction: column
	}

	.location-input__container:not(:focus-within) .possible-location-container {
		display: none;
	}

	.possible-location-container--hide {
		visibility: hidden;
		pointer-events: none;
	}

	.possible-location-container button {
		text-align: left;
		width: 250px;
		padding: 10px;
		height: auto;
	}


	.autoCompleteError {
		position: relative;
	}
	
	.autoCompleteError .error-tooltip {
		display: block;
	}

	.header {
		color: white;
		padding: 10px;
		text-align: center;
	}

	.header h2 {
		margin-bottom: 0;
	}
	
	.link-container {
		margin-bottom: 20px;
	}

	.link-container .link {
		padding: 5px 10px;
		background: white;
		color: #c0392b;
		cursor: pointer;
	}

	.instructions {
		display: flex;
		flex-wrap: wrap;
		width: 650px;
		color: white;
		line-height: 18px;
		margin-top: 10px;
		margin-bottom: 50px;
	}

	.instructions--hide {
		visibility: hidden;
	}
	
	.instructions__left {
    	width: calc(50% - 20px);
		padding: 0 10px 0 0;
	}

	.instructions__right {
    	width: calc(50% - 20px);
		padding: 0 0 0 10px;
	}

	@media screen and (max-width: 650px) {
		.instructions {
			width: auto;
		}
		.instructions__left{
			width: 100%;
			padding: 10px;
		}
		 .instructions__right {
			 width: 100%;
			padding: 10px;
		 }
	}

</style>

{#if flashError}
	<div class="flash-message">
		{ flashError }
	</div>
{/if}

<div class="container">
	<div class="header">
		<h2>Google Location Faker</h2>
		<p>See local map results as if you were in the location.</p>
	</div>
	<div class="input-container">
		<input class="query" bind:value={query}>
		<div class="location location-input__container { autoCompleteError ? "autoCompleteError" : "" }">
			<input bind:value={location} on:keyup={handleLocationInputChange} >
				<div class="error-tooltip">
					No location found in our database - however you can still try it. Note the input is case sensitive.
				</div>
				<div class="possible-location-container">
					{#each possibleLocations as location, i}
						<button on:click={handlePossibleLocationSubmit} on:keyup={handleLocationInputChange} id={`possibleLocation--${i}`} name={location}>
							{location} 
						</button>
					{/each}
				</div>
		</div>
		<button class="submit" on:click={handleClick}>
			Get Results
		</button>
	</div>
	<div class="link-container"> 
		{#if googleLink}
			<a class="link" href={googleLink} target="_blank">Link here!</a>
		{/if}
	</div>
	<div class="instructions { displayPossibleLocations ? "instructions--hide" : null } ">
		<div class="instructions__left">
			<h2 style="margin-top: 0">How does it work?</h2>
			Google uses an algorithm, which is based on the length of the input location, to encode the location data of a search. <br /><br /> We can fake the output of this algorithm to trick Google to show us local listings. 
		</div>
		<div class="instructions__right">
			The location input features an optional location autocomplete. This is useful if you are searching for somewhere that has a duplicate name.<br /><br /> The list of autocomplete names is provided via Google & is case sensitive.
		</div>
	</div>
</div>
