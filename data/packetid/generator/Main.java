package me.tnze;

import it.unimi.dsi.fastutil.ints.Int2ObjectMap;
import net.minecraft.network.ConnectionProtocol;
import net.minecraft.network.protocol.Packet;
import net.minecraft.network.protocol.PacketFlow;

public class Main {

    public static void main(String[] args) throws Exception {
        handlePackets(ConnectionProtocol.LOGIN.getPacketsByIds(PacketFlow.CLIENTBOUND));
        System.out.println();
        handlePackets(ConnectionProtocol.LOGIN.getPacketsByIds(PacketFlow.SERVERBOUND));
        System.out.println();
        System.out.println();
        handlePackets(ConnectionProtocol.STATUS.getPacketsByIds(PacketFlow.CLIENTBOUND));
        System.out.println();
        handlePackets(ConnectionProtocol.STATUS.getPacketsByIds(PacketFlow.SERVERBOUND));
        System.out.println();
        System.out.println();
        handlePackets(ConnectionProtocol.PLAY.getPacketsByIds(PacketFlow.CLIENTBOUND));
        System.out.println();
        handlePackets(ConnectionProtocol.PLAY.getPacketsByIds(PacketFlow.SERVERBOUND));
    }

    private static void handlePackets(Int2ObjectMap<Class<? extends Packet<?>>> packets) {
        for (int i = 0; i < packets.size(); i++) {
            Class<? extends Packet<?>> c = packets.get(i);
            String className = c.getSimpleName();
            if (className.endsWith("Packet"))
                className = className.substring(0, className.length() - "Packet".length());
            else {
                String superClassName = c.getSuperclass().getSimpleName();
                if (superClassName.endsWith("Packet"))
                    className = superClassName.substring(0, superClassName.length() - "Packet".length()) + className;
            }

            System.out.println(className);

	    var fields = c.getDeclaredFields();
	    for (var field : fields) {
		var name = field.getName();
		var type = field.getType();
		var typeName = type.getName();

		// no packet parameters are all UPPER CASE
		if (name.chars().allMatch(ch -> Character.isUpperCase(ch) || !Character.isAlphabetic(ch)))
			continue;

		System.out.println("\t" + name + "\t" + typeName);
	    }
        }
    }
}
