<?php

echo "\nData Extractor\n";
echo "Reading Biodata from file 'CCCSP FUTA Final-Year Brethren Biodata.csv'... (Ensure the Google form document is saved in the same folder as this program with exactly this name: CCCSP FUTA Final-Year Brethren Biodata.csv)\n";

try {
    $source = file_get_contents("CCCSP FUTA Final-Year Brethren Biodata.csv");

    if (!file_exists("Personal")) {
        mkdir("Personal");
    }

    $members = explode('"\n"', $source);
    $title = explode('","', $members[0]);

    for ($m = 1; $m < count($members); $m++) {
        $member = $members[$m];
        $biodata = explode('","', $member);
        $name = ucwords(strtolower(trim($biodata[2]) . " " . trim($biodata[3])));
        $filePath = "Personal/" . $name . ".txt";
        $output = fopen($filePath, "w");

        $line = "Name:\n" . $name . "\n\n\n";
        fwrite($output, $line);

        $line = "Email Address:\n" . trim($biodata[1]) . "\n\n\n";
        fwrite($output, $line);

        for ($n = 4; $n < count($biodata) - 1; $n++) {
            $entry = trim($biodata[$n]);
            if (strlen($entry) > 0) {
                $line = $title[$n] . ":\n" . $entry . "\n\n\n";
                fwrite($output, $line);
            }
        }

        fclose($output);
    }

    echo "SUCCESS: All data extracted successfully and saved in the folder 'Personal/'.";
    sleep(5);

} catch (Exception $e) {
    echo "FAILED: " . $e->getMessage();
    sleep(15);
}

?>
