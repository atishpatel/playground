n=2000
c=500
echo "Starting benchmark"
echo ""> bench.md
echo "# nginx -> go" >> bench.md
echo "\`\`\`" >> bench.md
hey -n=$n -c=$c http://localhost:9000/server1 >> bench.md
echo "\`\`\`" >> bench.md
echo "# go" >> bench.md
echo "\`\`\`" >> bench.md
hey -n=$n -c=$c http://localhost:9001/ >> bench.md
echo "\`\`\`" >> bench.md
echo "# nginx -> node" >> bench.md
echo "\`\`\`" >> bench.md
hey -n=$n -c=$c http://localhost:9000/server2 >> bench.md
echo "\`\`\`" >> bench.md
echo "# node" >> bench.md
echo "\`\`\`" >> bench.md
hey -n=$n -c=$c http://localhost:9002/ >> bench.md
echo "\`\`\`" >> bench.md