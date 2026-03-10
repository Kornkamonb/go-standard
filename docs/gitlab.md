git config --local user.name "smartfactory"
git config --local user.email "smartfactory@th.fujikura.com"
git config pull.rebase true
git remote add origin http://papop.p:glpat-XuRaFgjUvwGbfar0yAfgY286MQp1OmwH.01.0w1u8xb24@10.17.66.143:8005/llm/backend/backend-llm.git
git add .
git commit -m "$(hostname) update"
git fetch origin main
git pull origin main
git push --set-upstream origin main
