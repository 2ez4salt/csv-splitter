count=0
cd $1 && mkdir securedFiles
for d in *.csv
do
    count=$(( count + 1 ))
    secure-spreadsheet --password $2 < $d > ./securedFiles/$3_"$count".xlsx && echo "$d" dosya şifrelendi.
done
echo İşlem tamamlandı!