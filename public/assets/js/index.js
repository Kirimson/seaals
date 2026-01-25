let saysInput = document.getElementById("says-input");
let saysLink = document.getElementById("says-link");
// Reset input
saysInput.value = "Something";
UpdateSays();

getStats();

function UpdateSays() {
  saysLink.href = `/seal/says/${saysInput.value}`;
}

async function getStats() {
  const url = "/api/stats";
  let countText = document.getElementById("seal-count");
  let popularTags = document.getElementById("popular-tags");
  try {
    const response = await fetch(url);
    if (!response.ok) {
      throw new Error(`Response status: ${response.status}`);
    }

    const stats = await response.json();
    countText.innerHTML = stats["count"];
    // stats.tags.forEach((tag, idx) => {
    //   console.log(idx);
    //   row = popularTags.insertRow(idx);
    //   console.log(tag.name);
    //   nameCell = row.insertCell(0);
    //   nameCell.innerHTML = tag.name;
    //   console.log(tag.count);
    //   countCell = row.insertCell(1);
    //   countCell.innerHTML = tag.count;
    // });
  } catch (error) {
    console.error(error.message);
  }
}
