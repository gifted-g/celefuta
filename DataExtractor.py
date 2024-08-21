import os
import time

print("\nData Extractor\n")
print("Reading Biodata from file 'CCCSP FUTA Final-Year Brethren Biodata.csv'... (Ensure the Google form document is saved in the same folder as this program with exactly this name: CCCSP FUTA Final-Year Brethren Biodata.csv)\n")

try:
	source = open("CCCSP FUTA Final-Year Brethren Biodata.csv", "rb").read().decode("utf-8")

	try: os.mkdir("Personal")
	except: pass

	members = source.split('"\n"')

	title = members[0].split('","')

	for m in range(1, len(members)):    
		member = members[m]
		biodata = member.split('","')

		name = (biodata[2].strip()+" "+biodata[3].strip()).title()

		output = open("Personal/"+name+".txt", "wb")

		line = "Name:\n"+name+"\n\n\n"
		output.write(line.encode("utf-8"))

		line = "Email Address:\n"+biodata[1].strip()+"\n\n\n"
		output.write(line.encode("utf-8"))
		
		for n in range(4, len(biodata)-1):
			entry = biodata[n].strip()
			if len(entry) > 0:
				line = title[n]+":\n"+entry+"\n\n\n"
				output.write(line.encode("utf-8"))

		output.close()

	print("SUCCESS: All data extracted successfully and saved in the folder 'Personal/'.")
	time.sleep(5)


#Email *
olamitiwilson@gmail.com
Olamiti
Wilson Oyewole
DD
/
MM
Goodness
CCCSP FUTA Final-Year Brethren Biodata
Personal Information
The respondent's email (olamitiwilson@gmail.com) was recorded on submission of this form.
Surname *
Other Names *



except Exception as e:
	print(f"FAILED: {e}")
	time.sleep(15)
