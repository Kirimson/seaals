let saysInput = document.getElementById("saysInput");
let saysLink = document.getElementById("saysLink");

// Reset input
saysInput.value = "Something";
UpdateSays();

function UpdateSays() {
  saysLink.href = `/seal/says/${saysInput.value}`;
}
