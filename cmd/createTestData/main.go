package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/JesusJMM/monomio/postgres"
)

var ImgArr [15]string = [15]string{
  "https://images.unsplash.com/photo-1655552360620-5212a1b24961?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwzfHx8ZW58MHx8fHw%3D&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1531297484001-80022131f5a1?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8dGVjaHxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1655438819488-69bfb10a6c21?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw0fHx8ZW58MHx8fHw%3D&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1496065187959-7f07b8353c55?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MTV8fHRlY2h8ZW58MHx8MHx8&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1655372501819-4c1261a50c55?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwzMnx8fGVufDB8fHx8&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1518770660439-4636190af475?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MTB8fHRlY2h8ZW58MHx8MHx8&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1655012325191-cbc22182fa9f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHx0b3BpYy1mZWVkfDl8Q0R3dXdYSkFiRXd8fGVufDB8fHx8&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1653669718797-5670d0b57ca2?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHx0b3BpYy1mZWVkfDM5fENEd3V3WEpBYkV3fHxlbnwwfHx8fA%3D%3D&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1488590528505-98d2b5aba04b?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Mnx8dGVjaHxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1526228653958-2fda2c7479eb?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHx0b3BpYy1mZWVkfDJ8NnNNVmpUTFNrZVF8fGVufDB8fHx8&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1655212874354-5dace1fdc6ce?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHx0b3BpYy1mZWVkfDV8NnNNVmpUTFNrZVF8fGVufDB8fHx8&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1638913975386-d61f0ec6500d?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHwxMXx8fGVufDB8fHx8&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1655566655625-4aadace23511?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwzNXx8fGVufDB8fHx8&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1655493623673-13f063cb78cf?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw2NXx8fGVufDB8fHx8&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1655480530311-b3c63da9021e?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw3MXx8fGVufDB8fHx8&auto=format&fit=crop&w=200&q=60",
}

var BigImgArr [15]string = [15]string{
  "https://images.unsplash.com/photo-1655552360620-5212a1b24961?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwzfHx8ZW58MHx8fHw%3D&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1531297484001-80022131f5a1?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8dGVjaHxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1655438819488-69bfb10a6c21?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw0fHx8ZW58MHx8fHw%3D&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1496065187959-7f07b8353c55?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MTV8fHRlY2h8ZW58MHx8MHx8&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1655372501819-4c1261a50c55?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwzMnx8fGVufDB8fHx8&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1518770660439-4636190af475?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MTB8fHRlY2h8ZW58MHx8MHx8&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1655012325191-cbc22182fa9f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHx0b3BpYy1mZWVkfDl8Q0R3dXdYSkFiRXd8fGVufDB8fHx8&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1653669718797-5670d0b57ca2?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHx0b3BpYy1mZWVkfDM5fENEd3V3WEpBYkV3fHxlbnwwfHx8fA%3D%3D&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1488590528505-98d2b5aba04b?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Mnx8dGVjaHxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1526228653958-2fda2c7479eb?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHx0b3BpYy1mZWVkfDJ8NnNNVmpUTFNrZVF8fGVufDB8fHx8&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1655212874354-5dace1fdc6ce?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHx0b3BpYy1mZWVkfDV8NnNNVmpUTFNrZVF8fGVufDB8fHx8&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1638913975386-d61f0ec6500d?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHwxMXx8fGVufDB8fHx8&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1655566655625-4aadace23511?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwzNXx8fGVufDB8fHx8&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1655493623673-13f063cb78cf?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw2NXx8fGVufDB8fHx8&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1655480530311-b3c63da9021e?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw3MXx8fGVufDB8fHx8&auto=format&fit=crop&w=600&q=60",
}

func main() {
  dbConn, err := sql.Open("postgres", "user=monomioapp dbname=monomioapp sslmode=disable") 
  if err != nil {
    log.Fatalf("Error opening the database, %s", err.Error())
  }
  db := postgres.New(dbConn)
  // Clean up last test
  ctx := context.Background()
  old_user, err := db.GetUserByName(ctx, "Test User")
  if err != nil {
    log.Fatalf("Error opening the database, %s", err.Error())
  }
  dbConn.Exec(`DELETE FROM post WHERE user_id = ?`, old_user.ID)
  fmt.Println("Old test articles deleted")
  db.DeleteUser(ctx, old_user.ID)
  fmt.Println("Old test articles deleted")

  // Create User
  user, err := db.CreateUser(ctx, postgres.CreateUserParams{
    Name: "Test User",
    Password: "test_password",
    ImgUrl: sql.NullString{String: "https://images.unsplash.com/photo-1571771894821-ce9b6c11b08e?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=80&q=80", Valid: true},
  })
  if err != nil {
    log.Fatalf("Error creating user, %s", err.Error())
  }
  // Create Articles
  for i := 0; i < len(ImgArr); i ++ {
    article, err := db.CreatePost(ctx, postgres.CreatePostParams{
      UserID: user.ID,
      Title: "Article Title For tests",
      Description: sql.NullString{String: "Article for test propouse, if this article exist outside of test scope, must be deleted", Valid: true},
      Content: sql.NullString{String: AritcleBody, Valid: true},
      Slug: fmt.Sprintf("test-article-%d", i),
      FeedImg: sql.NullString{String: ImgArr[i], Valid: true},
      ArticleImg: sql.NullString{String: BigImgArr[i], Valid: true},
    })
    if err != nil {
      log.Fatalf("Error creating article #%d, %s",i , err.Error())
    }
    err = db.PublishPost(ctx, postgres.PublishPostParams{
      ID: article.ID,
      UserID: user.ID,
    })
    if err != nil {
      log.Fatalf("Error publishgin article #%d, %s",i , err.Error())
    }
    fmt.Printf("Article #%d created!\n", i)
  }
}


const AritcleBody = `# Deno

[![Build Status - Cirrus][]][Build status] [![Twitter handle][]][Twitter badge]
[![Discord Chat](https://img.shields.io/discord/684898665143206084?logo=discord&style=social)](https://discord.gg/deno)

<img align="right" src="https://deno.land/logo.svg" height="150px" alt="the deno mascot dinosaur standing in the rain">

Deno is a _simple_, _modern_ and _secure_ runtime for **JavaScript** and
**TypeScript** that uses V8 and is built in Rust.

### Features

- Secure by default. No file, network, or environment access, unless explicitly
  enabled.
- Supports TypeScript out of the box.
- Ships only a single executable file.
- [Built-in utilities.](https://deno.land/manual/tools#built-in-tooling)
- Set of reviewed standard modules that are guaranteed to work with
  [Deno](https://deno.land/std/).

### Install

Shell (Mac, Linux):

"""sh
curl -fsSL https://deno.land/install.sh | sh
"""

PowerShell (Windows):

"""powershell
iwr https://deno.land/install.ps1 -useb | iex
"""

[Homebrew](https://formulae.brew.sh/formula/deno) (Mac):

"""sh
brew install deno
"""

[Chocolatey](https://chocolatey.org/packages/deno) (Windows):

"""powershell
choco install deno
"""

[Scoop](https://scoop.sh/) (Windows):

"""powershell
scoop install deno
"""

Build and install from source using [Cargo](https://crates.io/crates/deno):

"""sh
cargo install deno --locked
"""

See
[deno_install](https://github.com/denoland/deno_install/blob/master/README.md)
and [releases](https://github.com/denoland/deno/releases) for other options.

### Getting Started

Try running a simple program:

"""sh
deno run https://deno.land/std/examples/welcome.ts
"""

Or a more complex one:

"""ts
const listener = Deno.listen({ port: 8000 });
console.log("http://localhost:8000/");

for await (const conn of listener) {
  serve(conn);
}

async function serve(conn: Deno.Conn) {
  for await (const { respondWith } of Deno.serveHttp(conn)) {
    respondWith(new Response("Hello world"));
  }
}
"""

You can find a deeper introduction, examples, and environment setup guides in
the [manual](https://deno.land/manual).

The complete API reference is available at the runtime
[documentation](https://doc.deno.land).

### Contributing

We appreciate your help!

To contribute, please read our
[contributing instructions](https://deno.land/manual/contributing).

[Build Status - Cirrus]: https://github.com/denoland/deno/workflows/ci/badge.svg?branch=main&event=push
[Build status]: https://github.com/denoland/deno/actions
[Twitter badge]: https://twitter.com/intent/follow?screen_name=deno_land
[Twitter handle]: https://img.shields.io/twitter/follow/deno_land.svg?style=social&label=Follow`
