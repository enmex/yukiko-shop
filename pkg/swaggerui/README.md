# Create static directory

## DEV Server
```bash
mkdir -p static
cp swagger.yaml static/swagger.yaml
cp ChangeLog static/ChangeLog
echo "User-agent: *" > static/robots.txt
echo "Disallow: /spec" >> static/robots.txt
cp pkg/swaggerui/static/index.html static/index.html
sed -ie 's/{{ path }}/\/nfs-open\/spec\/swagger.yaml/g' "static/index.html"
```

## Local Server
```bash
mkdir -p static
cp swagger.yaml static/swagger.yaml
cp ChangeLog static/ChangeLog
echo "User-agent: *" > static/robots.txt
echo "Disallow: /spec" >> static/robots.txt
cp pkg/swaggerui/static/index.html static/index.html
sed -ie 's/{{ path }}/\/spec\/swagger.yaml/g' "static/index.html"
```