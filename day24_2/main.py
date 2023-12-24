from z3 import Int, Solver

def parse_line(line : str):
    parts = line.split(' @ ')

    startPositions = [int(x.strip()) for x in parts[0].split(', ')]

    velocities = [int(x.strip()) for x in parts[1].split(', ')]


    return startPositions, velocities 


def main():
    with open('../data/input24.txt' ,'r') as f:
        content = f.read().splitlines()

    hailstones = [parse_line(line) for line in content] 
    
    solver = Solver()
    result_x = Int("result_x")
    result_y = Int("result_y")
    result_z = Int("result_z")
    result_velocity_x = Int("result_velocity_x")
    result_velocity_y = Int("result_velocity_y")
    result_velocity_z = Int("result_velocity_z")

    for i, ([x,y,z], [dx,dy,dz]) in enumerate(hailstones):
        t = Int(f"t{i}")
        solver.add(t >= 0)
        solver.add(x + dx * t == result_x + result_velocity_x * t)
        solver.add(y + dy * t == result_y + result_velocity_y * t)
        solver.add(z + dz * t == result_z + result_velocity_z * t)

        if i == 3:
            break

    assert str(solver.check()) == 'sat'
    print( solver.model().eval(result_x + result_y + result_z))



if __name__ == '__main__':
    main()
