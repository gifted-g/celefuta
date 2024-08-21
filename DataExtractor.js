const fs = require('fs');
const path = require('path');

console.log("\nData Extractor\n");
console.log("Reading Biodata from file 'CCCSP FUTA Final-Year Brethren Biodata.csv'... (Ensure the Google form document is saved in the same folder as this program with exactly this name: CCCSP FUTA Final-Year Brethren Biodata.csv)\n");

try {
    const source = fs.readFileSync("CCCSP FUTA Final-Year Brethren Biodata.csv", "utf8");

    // Create the Personal folder if it doesn't exist
    if (!fs.existsSync("Personal")) {
        fs.mkdirSync("Personal");
    }

    const members = source.split('"\n"');
    const title = members[0].split('","');

    for (let m = 1; m < members.length; m++) {
        const member = members[m];
        const biodata = member.split('","');
        const name = (biodata[2].trim() + " " + biodata[3].trim()).toLowerCase().replace(/\b\w/g, l => l.toUpperCase());
        const filePath = path.join("Personal", `${name}.txt`);
        const output = [];

        output.push(`Name:\n${name}\n\n\n`);
        output.push(`Email Address:\n${biodata[1].trim()}\n\n\n`);

        for (let n = 4; n < biodata.length - 1; n++) {
            const entry = biodata[n].trim();
            if (entry.length > 0) {
                output.push(`${title[n]}:\n${entry}\n\n\n`);
            }
        }

        fs.writeFileSync(filePath, output.join(""), "utf8");
    }

    console.log("SUCCESS: All data extracted successfully and saved in the folder 'Personal/'.");
    setTimeout(() => {}, 5000);

} catch (e) {
    console.log(`FAILED: ${e}`);
    setTimeout(() => {}, 15000);
}
