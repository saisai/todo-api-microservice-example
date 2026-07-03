```
…or create a new repository on the command line

echo "# todo-api-microservice-example" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/saisai/todo-api-microservice-example.git
git push -u origin main
```

```
…or push an existing repository from the command line

git remote add origin https://github.com/saisai/todo-api-microservice-example.git
git branch -M main
git push -u origin main
```

# To get the oldest (initial) commit in a Git repository, use the git log --reverse command combined with a tool to limit the output.

```
git log --reverse --oneline | head -n 5

```

```
git log --reverse --format="%H" | head -n 5
git log --reverse | head -n 25
```

```
git log --reverse --format="%H %s" | head -n 5
```

```
git log --reverse --format="%H" | head -n 5
```