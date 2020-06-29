package base;

class Philosopher implements Runnable {
    private final String name;
    private final Object leftChopstick;
    private final Object rightChopstick;

    Philosopher(String name, Object left, Object right) {
        this.name = name;
        leftChopstick = left;
        rightChopstick = right;
    }

    @Override
    public void run() {
        while (true) {
            System.out.println(System.nanoTime() + "\t" + name + ": Thinking");
            synchronized (leftChopstick) {
                synchronized (rightChopstick) {
                    System.out.println(System.nanoTime() + "\t" + name + ": Eating");
                }
            }
        }
    }
}

public class DiningPhilosophers {
    static final int COUNT = 5;

    public static void main(String[] args) throws InterruptedException {
        Object[] chopsticks = new Object[COUNT];
        for (int i = 0; i < COUNT; i++) {
            chopsticks[i] = new Object();
        }

        for (int i = 0; i < COUNT - 1; i++) {
            new Thread(new Philosopher("Number " + i, chopsticks[i + 1], chopsticks[i])).start();
        }
        new Thread(new Philosopher("Number " + (COUNT - 1), chopsticks[COUNT - 1], chopsticks[0])).start();

        Thread.sleep(1000);
        System.exit(0);
    }
}
