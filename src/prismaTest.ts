import { PrismaClient } from '@prisma/client'

const prisma = new PrismaClient()

async function main() {

  const tagName = "test"
  // Get the tag
  let testTag = await prisma.tag.findUnique({
    where: {
      name: tagName
    }
  })

  if (!testTag) {
    testTag = await prisma.tag.create({
      data: {
        name: tagName
      }
    })
  }

  const sealName = 'Alice'
  let seal = await prisma.seal.findUnique({
    where: {
      slug: sealName
    }
  })
  // Create seal if not exists
  if (!seal) {
    seal = await prisma.seal.create({
      data: {
        slug: sealName,
        tags: {
          connect: [{id: testTag.id}]
        },
      },
    })
  }

  // Get seal and its tags
  let sealAndTag = await prisma.seal.findUnique({
    where: {
      slug: sealName
    },
    include: {
      tags: true
    }
  })
  console.log(sealAndTag)
}

main()

  .then(async () => {
    await prisma.$disconnect()
  })
  .catch(async (e) => {
    console.error(e)
    await prisma.$disconnect()
    process.exit(1)

  })