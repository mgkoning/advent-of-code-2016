// Learn more about F# at http://fsharp.org
// See the 'F# Tutorial' project for more help.
open Microsoft.FSharp.Math
open System

type Coordinate = {
  x: int
  y: int
}

type Facing = North | East | South | West

type Position = {
  coordinate: Coordinate
  facing: Facing
}

type Turn = Left | Right

type Direction = {
  turn: Turn
  blocks: int
}

let distanceFromOrigin coordinate =
  abs coordinate.x + abs coordinate.y

let sampleInput1 = "R2, L3"
let sampleInput2 = "R2, R2, R2"
let sampleInput3 = "R5, L5, R5, R3"
let puzzleInput1 = "L5, R1, L5, L1, R5, R1, R1, L4, L1, L3, R2, R4, L4, L1, L1, R2, R4, R3, L1, R4, L4, L5, L4, R4, L5, R1, R5, L2, R1, R3, L2, L4, L4, R1, L192, R5, R1, R4, L5, L4, R5, L1, L1, R48, R5, R5, L2, R4, R4, R1, R3, L1, L4, L5, R1, L4, L2, L5, R5, L2, R74, R4, L1, R188, R5, L4, L2, R5, R2, L4, R4, R3, R3, R2, R1, L3, L2, L5, L5, L2, L1, R1, R5, R4, L3, R5, L1, L3, R4, L1, L3, L2, R1, R3, R2, R5, L3, L1, L1, R5, L4, L5, R5, R2, L5, R2, L1, L5, L3, L5, L5, L1, R1, L4, L3, L1, R2, R5, L1, L3, R4, R5, L4, L1, R5, L1, R5, R5, R5, R2, R1, R2, L5, L5, L5, R4, L5, L4, L4, R5, L2, R1, R5, L1, L5, R4, L3, R4, L2, R3, R3, R3, L2, L2, L2, L1, L4, R3, L4, L2, R2, R5, L1, R2"

let parseDirection (direction: string) =
  let leftOrRight = direction.[0]
  let blocks = Int32.Parse(direction.Substring(1))
  match leftOrRight with
  | 'L' -> { turn = Left; blocks = blocks }
  | 'R' -> { turn = Right; blocks = blocks }
  | _ -> failwith (sprintf "Direction %A not understood." leftOrRight)

let getDirections (input: string) =
  input.Split([|", "|], System.StringSplitOptions.None)
  |> Array.map parseDirection

let newFacing current direction =
  match (current, direction) with
  | North, Left -> West
  | East, Left -> North
  | South, Left -> East
  | West, Left -> South
  | North, Right -> East
  | East, Right -> South
  | South, Right -> West
  | West, Right -> North

let factorX = function West -> -1 | East -> 1 | _ -> 0
let factorY = function North -> 1 | South -> -1 | _ -> 0

let newCoordinate current facing blocks =
  { x = current.x + (factorX facing) * blocks
    y = current.y + (factorY facing) * blocks }

let newPosition (current: Position) direction =
  let newFacing = newFacing current.facing direction.turn
  { facing = newFacing;
    coordinate = newCoordinate current.coordinate newFacing direction.blocks }

let normalize i = i / (abs i)

let locations current facing blocks =
  let factorX = factorX facing
  let factorY = factorY facing
  ([1 .. (abs factorX) * blocks] |> List.map (fun i -> { x = current.x + (normalize factorX) * i; y = current.y }))
  @ ([1 .. (abs factorY) * blocks] |> List.map (fun i -> { x = current.x; y = current.y + (normalize factorY) * i }))


let initialPosition = {
  coordinate = { x = 0; y = 0 }
  facing = North
}

let finalCoordinate directions =
  Array.fold newPosition initialPosition directions 

let allPositions directions =
  let rec allPositionsHelper current directions allPositions =
    match current, directions with
    | _, [] -> allPositions
    | _, direction::rest ->
        let nextPosition = (newPosition current direction)
        let locations = (locations current.coordinate nextPosition.facing direction.blocks)
        allPositionsHelper nextPosition rest (allPositions@locations)
  allPositionsHelper initialPosition directions [initialPosition.coordinate]

let rec findFirstDuplicate (candidates: Coordinate list) (alreadySeen: Coordinate list) =
  match candidates with
  | [] -> failwith "no duplicates found"
  | candidate::rest when (List.exists (fun (c: Coordinate) -> c = candidate) alreadySeen) -> candidate
  | candidate::rest -> findFirstDuplicate rest (candidate::alreadySeen)

let solution input =
  distanceFromOrigin ((getDirections input) |> finalCoordinate).coordinate


[<EntryPoint>]
let main argv = 
    printfn "%A" (solution sampleInput1)
    printfn "%A" (solution sampleInput2)
    printfn "%A" (solution sampleInput3)
    printfn "%A" (solution puzzleInput1)
    let allPositions = allPositions (List.ofArray (getDirections puzzleInput1))
    let firstDuplicate = findFirstDuplicate allPositions []
    printfn "%A" firstDuplicate
    printfn "%A" (distanceFromOrigin firstDuplicate)
    0 // return an integer exit code
